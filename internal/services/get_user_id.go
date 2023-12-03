package services

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"uzum_shop/generated/protos/login_v1"
)

func (s *ShopService) GetUserId(ctx context.Context) (uuid.UUID, error) {
	emp := &login_v1.GetData_Request{EndpointAddress: ""}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	res, err := s.loginClient.GetData(ctx, emp)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(res.UserId), err
}
