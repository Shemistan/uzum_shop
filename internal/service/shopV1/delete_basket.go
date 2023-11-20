package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (s *shopSystemService) DeleteBasketService(ctx context.Context, req *pb.DeleteBasket_Request) error {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return err
	}

	deleteReq := &models.DeleteFomBasked{
		UserId:    userId,
		ProductId: req.ProductId,
	}
	err = s.storage.DeleteBasketStorage(ctx, deleteReq)
	if err != nil {
		return err
	}

	return nil
}
