package shop_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Shemistan/uzum_shop/internal/converter"
	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"
)

func (s *Shop) GetAllProduct(ctx context.Context, _ *emptypb.Empty) (*pb.GetProducts_Response, error) {
	res, err := s.ShopService.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]*pb.ShortProduct, 0, len(res))

	for _, v := range res {
		products = append(products, converter.ShortProductApiToModelService(v))
	}

	return &pb.GetProducts_Response{
		ShortProduct: products,
	}, nil
}
