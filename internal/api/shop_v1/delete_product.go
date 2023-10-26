package shop_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"
)

func (s *Shop) DeleteProduct(ctx context.Context, req *pb.DeleteProduct_Request) (*emptypb.Empty, error) {
	err := s.ShopService.DeleteProduct(ctx, req.GetProductId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
