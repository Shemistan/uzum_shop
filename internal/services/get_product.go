package services

import (
	"context"
	"github.com/google/uuid"
	"uzum_shop/internal/models"
)

func (s *ShopService) GetProduct(ctx context.Context, id uuid.UUID) (models.Product, error) {
	product, err := s.storage.GetProduct(ctx, id)
	if err != nil {
		return models.Product{}, err
	}
	return models.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description.String,
		Price:       product.Price.Float64,
		Quantity:    product.Quantity,
	}, err
}
