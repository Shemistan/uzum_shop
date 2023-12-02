package api

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/pkg/shopV1"
)

func (s *ShopAPI) CancelOrder(ctx context.Context, req *shopV1.Order_CancelRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.service.CancelOrder(ctx, uuid.MustParse(req.OrderId))
}
