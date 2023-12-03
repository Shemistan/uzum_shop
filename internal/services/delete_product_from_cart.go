package services

import (
	"context"
	"github.com/google/uuid"
)

func (s *ShopService) DeleteProductFromCart(ctx context.Context, productID uuid.UUID) error {
	userID, err := s.GetUserId(ctx)
	if err != nil {
		return err
	}

	delete(s.cart[userID].Items, productID)

	return nil
}
