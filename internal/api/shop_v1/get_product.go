package shop_v1

import (
	"context"

	"github.com/Shemistan/uzum_shop/internal/converter"
	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"
)

func (s *Shop) GetProduct(ctx context.Context, req *pb.GetProduct_Request) (*pb.GetProduct_Response, error) {
	res, err := s.ShopService.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetProduct_Response{Product: converter.ProductApiToModelService(res)}, nil
}
