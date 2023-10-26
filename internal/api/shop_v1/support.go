package shop_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Shop) Healthz(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
