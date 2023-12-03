package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s *shopSystemService) GetProductService(ctx context.Context, prodId uint32) (*models.Product, error) {
	response, err := s.storage.GetProductStorage(ctx, prodId)
	if err != nil {
		return nil, err
	}

	return response, nil
}
