package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"uzum_shop/generated/protos/login_v1"
	"uzum_shop/internal/api"
	"uzum_shop/internal/models"
)

type App struct {
	dbConn *sqlx.DB

	shopAPI     api.ShopAPI
	loginClient login_v1.LoginV1Client
	grpcServer  *grpc.Server
	httpMux     *runtime.ServeMux

	config *models.Config
}

func NewApp(ctx context.Context) (*App, error) {
	var app App
	err := app.loadConfig()
	if err != nil {
		return nil, err
	}
	err = app.initAPI()
	if err != nil {
		return nil, err
	}
	app.InitGRPC()
	err = app.InitHTTP(ctx)
	if err != nil {
		return nil, err
	}

	return &app, err
}
