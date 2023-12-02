package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/pkg/shopV1"
)

func (s *ShopAPI) GetProducts(ctx context.Context, req *emptypb.Empty) (*shopV1.ProductsGet_Response, error) {
	products, err := s.service.GetProducts(ctx)
	if err != nil {
		return &shopV1.ProductsGet_Response{}, err
	}

	var response []*shopV1.ShortProductInfo

	for _, product := range products {
		response = append(response, &shopV1.ShortProductInfo{ProductId: product.ID.String(),
			Name:  product.Name,
			Price: float32(product.Price),
		})
	}
	return &shopV1.ProductsGet_Response{Products: response}, nil
}