package api

import (
	"uzum_shop/internal/services"
	"uzum_shop/pkg/shopV1"
)

type ShopAPI struct {
	shopV1.UnimplementedShopServiceServer
	service *services.ShopService
}

func NewShopAPI(service *services.ShopService) ShopAPI {
	return ShopAPI{
		service: service,
	}
}
