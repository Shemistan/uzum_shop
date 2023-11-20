package shopV1

import (
	"context"
	"fmt"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (s *shopSystemService) UpdateBasketService(ctx context.Context, req *pb.UpdateBasket_Request) error {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return err
	}

	updateBask := &models.AddProductToBasketModel{
		UserId:    userId,
		ProductId: req.ProductId,
		Count:     req.Count,
	}
	res, err := s.storage.UpdateBasketStorage(ctx, updateBask)
	if err != nil {
		return err
	}

	if res == 0 {
		return fmt.Errorf("%v rows affected", res)
	}

	return nil
}
