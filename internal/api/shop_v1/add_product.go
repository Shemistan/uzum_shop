package shop_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Shemistan/uzum_shop/internal/converter"
	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"
)

func (s *Shop) AddProduct(ctx context.Context, req *pb.AddProduct_Request) (*emptypb.Empty, error) {
	err := s.ShopService.AddProduct(ctx, converter.AddProductRequestApiToModelService(req.GetBasket()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
