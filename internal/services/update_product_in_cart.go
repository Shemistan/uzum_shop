package services

import (
	"context"
	"github.com/google/uuid"
)

func (s *ShopService) UpdateProductInCart(ctx context.Context, productID uuid.UUID, quantity int32) error {
	userID, err := s.GetUserId(ctx)
	if err != nil {
		return err
	}

	productPrice, err := s.storage.GetProductPrice(ctx, productID)
	if err != nil {
		return err
	}
	diff := s.cart[userID].Items[productID] - quantity
	s.cart[userID].Items[productID] = quantity
	s.cart[userID].Total += float64(diff) * productPrice.Float64

	return nil
}
