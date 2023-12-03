package services

import (
	"context"
	"github.com/google/uuid"
	"uzum_shop/internal/models"
)

func (s *ShopService) AddProductToCart(ctx context.Context, productID uuid.UUID, quantity int32) error {
	userID, err := s.GetUserId(ctx)
	if err != nil {
		return err
	}

	if s.cart[userID] == nil {
		s.cart[userID] = models.NewCart()
	}

	productPrice, err := s.storage.GetProductPrice(ctx, productID)

	s.cart[userID].Items[productID] = quantity
	s.cart[userID].Total += productPrice.Float64 * float64(s.cart[userID].Items[productID])

	return nil
}
