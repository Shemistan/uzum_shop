package storage

import (
	"context"

	"github.com/Shemistan/uzum_shop/internal/models"
)

type IStorage interface {
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	AddProduct(ctx context.Context, req *models.Basket) error
	UpdateBasket(ctx context.Context, req *models.UpdateBasketRequest) error
	DeleteProduct(ctx context.Context, id int64) error
	GetBasket(ctx context.Context) ([]*models.Basket, error)
}
