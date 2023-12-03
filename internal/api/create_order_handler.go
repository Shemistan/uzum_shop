package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/generated/protos/shop_v1"
	"uzum_shop/internal/models"
)

func (s *ShopAPI) CreateOrder(ctx context.Context, req *shop_v1.Order_CreateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.CreateOrder(ctx, &models.Coordinates{
		Latitude:  float64(req.Latitude),
		Longitude: float64(req.Longitude),
	}, req.Metadata)
}
