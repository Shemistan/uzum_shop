package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) AddProductToBasket(ctx context.Context, req *pb.AddProductToBasket_Request) (*empty.Empty, error) {
	err := serv.ShopService.AddProductToBasketService(ctx, req)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
