package service

import (
	"context"

	"github.com/Shemistan/uzum_shop/internal/models"
	repo "github.com/Shemistan/uzum_shop/internal/storage"
)

type IService interface {
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	AddProduct(ctx context.Context, req *models.Basket) error
	UpdateBasket(ctx context.Context, req *models.UpdateBasketRequest) error
	DeleteProduct(ctx context.Context, id int64) error
	GetBasket(ctx context.Context) ([]*models.Basket, error)
}

func NewService(appConfig *models.Config, storage repo.IStorage) IService {
	return &service{
		appConfig: appConfig,
		storage:   storage,
	}
}

type service struct {
	appConfig *models.Config
	storage   repo.IStorage
}

func (s service) GetBasket(ctx context.Context) ([]*models.Basket, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteProduct(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s service) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
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
