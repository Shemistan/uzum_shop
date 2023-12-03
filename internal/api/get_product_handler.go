package api

import (
	"context"
	"github.com/google/uuid"
	"uzum_shop/generated/protos/shop_v1"
)

func (s *ShopAPI) GetProduct(ctx context.Context, req *shop_v1.ProductGet_Request) (*shop_v1.ProductGet_Response, error) {
	product, err := s.service.GetProduct(ctx, uuid.MustParse(req.ProductId))
	if err != nil {
		return &shop_v1.ProductGet_Response{}, err
	}
	response := &shop_v1.ProductGet_Response{
		Product: &shop_v1.Product{ProductId: product.ID.String(),
			Name:        product.Name,
			Price:       float32(product.Price),
			Description: product.Description,
			Quantity:    uint32(product.Quantity)},
	}

	return response, nil
}
