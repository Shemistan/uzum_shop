package service

import (
	"context"

	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s service) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	// check user
	_, err := s.userId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.storage.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
