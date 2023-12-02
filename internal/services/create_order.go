package services

import (
	"context"
	"github.com/google/uuid"
	"uzum_shop/internal/db"
	"uzum_shop/internal/models"
)

func (s *ShopService) CreateOrder(ctx context.Context, coordinates *models.Coordinates, metadata string) error {
	userID, err := s.GetUserId(ctx)
	if err != nil {
		return err
	}

	var basketsID []uuid.UUID

	user, err := s.storage.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	for productID, quantity := range s.cart[userID].Items {
		basketID, err := s.storage.AddBasketItem(ctx, db.AddBasketItemParams{ProductID: productID, Quantity: quantity})
		if err != nil {
			return err
		}
		basketsID = append(basketsID, basketID)
	}

	err = s.storage.CreateOrder(ctx, db.CreateOrderParams{
		UserID:           userID,
		BasketsID:        basketsID,
		Address:          user.Address,
		AddressLocationX: user.LocationX,
		AddressLocationY: user.LocationY,
		PickUpLocationX:  coordinates.Latitude,
		PickUpLocationY:  coordinates.Longitude,
		Metadata:         metadata,
		TotalPrice:       s.cart[userID].Total,
	})
	if err != nil {
		return err
	}

	return nil
}
