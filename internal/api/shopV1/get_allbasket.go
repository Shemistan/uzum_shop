package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) GetBasket(ctx context.Context, _ *empty.Empty) (*pb.GetBasket_Response, error) {
	resp, err := serv.ShopService.GetBasketService(ctx)
	if err != nil {
		return &pb.GetBasket_Response{}, err
	}

	return resp, nil
}
