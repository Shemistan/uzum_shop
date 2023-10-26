package shop_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Shemistan/uzum_shop/internal/converter"
	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"
)

func (s *Shop) GetBasket(ctx context.Context, _ *emptypb.Empty) (*pb.GetBasket_Response, error) {
	res, err := s.ShopService.GetBasket(ctx)
	if err != nil {
		return nil, err
	}

	baskets := make([]*pb.Basket, 0, len(res))

	for _, v := range res {
		baskets = append(baskets, converter.BasketModelServiceToApi(v))
	}

	return &pb.GetBasket_Response{
		Basket: baskets,
	}, nil
}
