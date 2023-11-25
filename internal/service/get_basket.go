package service

import (
	"context"

	"github.com/Shemistan/uzum_shop/internal/models"
)

func (s service) GetBasket(ctx context.Context) ([]*models.Basket, error) {
	userId, err := s.userId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.storage.GetBasket(ctx, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
