package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (s *shopSystemService) GetBasketService(ctx context.Context) (*pb.GetBasket_Response, error) {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return &pb.GetBasket_Response{}, err
	}
	response, err := s.storage.GetBasketStorage(ctx, userId)
	if err != nil {
		return &pb.GetBasket_Response{}, err
	}

	return response, nil
}
