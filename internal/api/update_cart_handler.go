package api

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/generated/protos/shop_v1"
)

func (s *ShopAPI) UpdateCart(ctx context.Context, req *shop_v1.Cart_UpdateRequest) (*emptypb.Empty, error) {
	err := s.service.UpdateProductInCart(ctx, uuid.MustParse(req.ProductId), int32(req.Quantity))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
