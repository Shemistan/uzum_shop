package services

import (
	"context"
	"uzum_shop/internal/models"
)

func (s *ShopService) GetCart(ctx context.Context) (*[]models.Product, error) {
	userID, err := s.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	cart := s.cart[userID]

	var cartItems []models.Product

	for productID, quantity := range cart.Items {
		product, err := s.storage.GetProduct(ctx, productID)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, models.Product{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price.Float64,
			Quantity:    quantity,
			Description: product.Description.String,
		})
	}

	return &cartItems, nil
}
