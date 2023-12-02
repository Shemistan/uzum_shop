package services

import (
	"context"
	"uzum_shop/internal/models"
)

func (s *ShopService) GetProducts(ctx context.Context) ([]models.ShortProductInfo, error) {
	products, err := s.storage.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	var shortProducts []models.ShortProductInfo

	for _, product := range products {
		shortProducts = append(shortProducts, models.ShortProductInfo{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price.Float64,
		})
	}
	return shortProducts, nil
}
