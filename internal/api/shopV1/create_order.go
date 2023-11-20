package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (serv *Shop) CreateOrder(ctx context.Context, req *pb.Order_Request) (*pb.Order_Response, error) {
	resp, err := serv.ShopService.CreateOrderService(ctx, req)
	if err != nil {
		return &pb.Order_Response{}, err
	}
	return resp, nil
}
