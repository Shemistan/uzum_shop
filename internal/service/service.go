package service

import (
	"context"

	"github.com/Shemistan/uzum_shop/internal/models"
	repo "github.com/Shemistan/uzum_shop/internal/storage"
	"github.com/Shemistan/uzum_shop/pkg/login_v1"
)

type IService interface {
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	AddProduct(ctx context.Context, req *models.Basket) error
	UpdateBasket(ctx context.Context, req *models.UpdateBasketRequest) error
	DeleteProduct(ctx context.Context, id int64) error
	GetBasket(ctx context.Context) ([]*models.Basket, error)
}

func NewService(appConfig *models.Config, storage repo.IStorage, loginClient login_v1.LoginV1Client) IService {
	return &service{
		appConfig:   appConfig,
		storage:     storage,
		loginClient: loginClient,
	}
}

type service struct {
	appConfig   *models.Config
	storage     repo.IStorage
	loginClient login_v1.LoginV1Client
}

func (s service) userId(ctx context.Context) (string, error) {
	checkResp, err := s.loginClient.Check(ctx, &login_v1.Check_Request{EndpointAddress: ""})
	if err != nil {
		return "", err
	}
	return checkResp.UserId, nil
}

func (s service) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) AddProduct(ctx context.Context, req *models.Basket) error {
	//TODO implement me
	panic("implement me")
}

func (s service) UpdateBasket(ctx context.Context, req *models.UpdateBasketRequest) error {
	//TODO implement me
	panic("implement me")
}
