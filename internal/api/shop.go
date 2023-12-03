package api

import (
	"uzum_shop/generated/protos/shop_v1"
	"uzum_shop/internal/services"
)

type ShopAPI struct {
	shop_v1.UnimplementedShopServiceServer
	service *services.ShopService
}

func NewShopAPI(service *services.ShopService) ShopAPI {
	return ShopAPI{
		service: service,
	}
}
