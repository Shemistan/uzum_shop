package api

import (
	"context"
	"github.com/google/uuid"
	"uzum_shop/pkg/shopV1"
)

func (s *ShopAPI) GetProduct(ctx context.Context, req *shopV1.ProductGet_Request) (*shopV1.ProductGet_Response, error) {
	product, err := s.service.GetProduct(ctx, uuid.MustParse(req.ProductId))
	if err != nil {
		return &shopV1.ProductGet_Response{}, err
	}
	response := &shopV1.ProductGet_Response{
		Product: &shopV1.Product{ProductId: product.ID.String(),
			Name:        product.Name,
			Price:       float32(product.Price),
			Description: product.Description,
			Quantity:    uint32(product.Quantity)},
	}

	return response, nil
}
