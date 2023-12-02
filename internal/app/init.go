package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"uzum_shop/build"
	"uzum_shop/internal/api"
	"uzum_shop/internal/db"
	"uzum_shop/internal/models"
	"uzum_shop/internal/services"
	"uzum_shop/pkg/loginV1"
	"uzum_shop/pkg/shopV1"
	storage "uzum_shop/storage/postgres"
)

func (app *App) InitGRPC() {
	app.grpcServer = grpc.NewServer()

	shopV1.RegisterShopServiceServer(app.grpcServer, &app.shopAPI)
}

func (app *App) InitHTTP(ctx context.Context) error {
	app.httpMux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := shopV1.RegisterShopServiceHandlerFromEndpoint(ctx, app.httpMux, app.config.Host+app.config.GrpcPort, opts)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initAPI() (err error) {
	app.dbConn, err = storage.NewPostgres(app.config.DbConnStr)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", app.config.Host, app.config.LoginClientPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	app.loginClient = loginV1.NewLoginV1Client(conn)

	shopService := services.NewShopService(db.New(app.dbConn), app.loginClient)
	app.shopAPI = api.NewShopAPI(shopService)

	return err
}

func (app *App) loadConfig() error {
	var conf models.Config

	if build.DEBUG {
		err := build.SetConfig()
		if err != nil {
			return err
		}
	}
	err := envconfig.Process("", &conf)

	app.config = &conf
	return err
}
