package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

func (s *ShopService) CancelOrder(ctx context.Context, orderID uuid.UUID) error {
	userID, err := s.GetUserId(ctx)
	if err != nil {
		return err
	}

	orderOwner, err := s.storage.GetOrderOwner(ctx, orderID)
	if err != nil {
		return err
	}

	if orderOwner == userID {
		return s.storage.CancelOrder(ctx, orderID)
	}

	return errors.New("you can't cancel this order")
}
