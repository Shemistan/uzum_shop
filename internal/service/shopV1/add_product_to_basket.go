package shopV1

import (
	"context"
	"fmt"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (s *shopSystemService) AddProductToBasketService(ctx context.Context, req *pb.AddProductToBasket_Request) error {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return err
	}

	count, err := s.storage.GetProductCountStorage(ctx, req.ProductId)
	if err != nil {
		return err
	}

	if count < int(req.Count) {
		return fmt.Errorf("sorry, not enough inventory. Only %v left in stock", count)
	}

	addProdToBask := &models.AddProductToBasketModel{
		UserId:    userId,
		ProductId: req.ProductId,
		Count:     req.Count,
	}

	err = s.storage.AddProductToBasketStorage(ctx, addProdToBask)
	if err != nil {
		return err
	}

	count = count - int(req.Count)

	res, err := s.storage.CalculateProductCountStorage(ctx, int(req.ProductId), count)
	if err != nil {
		return err
	}
	if res == 0 {
		return fmt.Errorf("%v rows affected", res)
	}

	return nil
}
