package api

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/generated/protos/shop_v1"
)

func (s *ShopAPI) CancelOrder(ctx context.Context, req *shop_v1.Order_CancelRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.CancelOrder(ctx, uuid.MustParse(req.OrderId))
}
