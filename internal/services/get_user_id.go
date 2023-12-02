package services

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"uzum_shop/pkg/loginV1"
)

func (s *ShopService) GetUserId(ctx context.Context) (uuid.UUID, error) {
	emp := &loginV1.Check_Request{EndpointAddress: ""}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	res, err := s.loginClient.Check(ctx, emp)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(res.UserId), err
}
