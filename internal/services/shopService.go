package services

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"uzum_shop/generated/protos/login_v1"
	"uzum_shop/internal/db"
	"uzum_shop/internal/models"
)

type (
	ShopService struct {
		storage     storage
		loginClient login_v1.LoginV1Client
		cart        map[uuid.UUID]*models.Cart
	}
	storage interface {
		GetProduct(ctx context.Context, id uuid.UUID) (db.Product, error)
		GetProducts(ctx context.Context) ([]db.Product, error)
		CreateOrder(ctx context.Context, arg db.CreateOrderParams) error
		GetUser(ctx context.Context, id uuid.UUID) (db.GetUserRow, error)
		GetProductPrice(ctx context.Context, id uuid.UUID) (sql.NullFloat64, error)
		AddBasketItem(ctx context.Context, arg db.AddBasketItemParams) (uuid.UUID, error)
		DeleteBasketItem(ctx context.Context, id uuid.UUID) error
		GetOrderOwner(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
		CancelOrder(ctx context.Context, id uuid.UUID) error
	}
)

func NewShopService(storage storage, loginClient login_v1.LoginV1Client) *ShopService {
	return &ShopService{storage: storage,
		cart:        make(map[uuid.UUID]*models.Cart),
		loginClient: loginClient}
}
