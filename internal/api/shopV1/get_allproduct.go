package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (serv *Shop) GetAllProduct(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllProducts_Response, error) {
	resp, err := serv.ShopService.GetAllProductsService(ctx)
	if err != nil {
		return &pb.GetAllProducts_Response{}, err
	}

	return resp, nil
}
