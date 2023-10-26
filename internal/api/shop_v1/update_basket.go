package shop_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"

	"github.com/Shemistan/uzum_shop/internal/converter"
)

func (s *Shop) UpdateBasket(ctx context.Context, req *pb.UpdateBasket_Request) (*emptypb.Empty, error) {
	err := s.ShopService.UpdateBasket(ctx, converter.UpdateBasketRequestApiToModelService(req.GetUpBasket()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
