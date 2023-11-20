package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (serv *Shop) GetProduct(ctx context.Context, req *pb.GetProduct_Request) (*pb.GetProduct_Response, error) {
	resp, err := serv.ShopService.GetProductService(ctx, req)
	if err != nil {
		return &pb.GetProduct_Response{}, err
	}

	return resp, nil
}
