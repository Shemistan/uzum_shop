package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"uzum_shop/build"
	"uzum_shop/generated/protos/login_v1"
	"uzum_shop/generated/protos/shop_v1"
	"uzum_shop/internal/api"
	"uzum_shop/internal/db"
	"uzum_shop/internal/models"
	"uzum_shop/internal/services"
	storage "uzum_shop/storage/postgres"
)

func (app *App) InitGRPC() {
	app.grpcServer = grpc.NewServer()

	shop_v1.RegisterShopServiceServer(app.grpcServer, &app.shopAPI)
}

func (app *App) InitHTTP(ctx context.Context) error {
	app.httpMux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := shop_v1.RegisterShopServiceHandlerFromEndpoint(ctx, app.httpMux, app.config.Host+app.config.GrpcPort, opts)
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

	conn, err := grpc.Dial(app.config.Host+app.config.LoginClientPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	app.loginClient = login_v1.NewLoginV1Client(conn)

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
