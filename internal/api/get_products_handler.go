package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/generated/protos/shop_v1"
)

func (s *ShopAPI) GetProducts(ctx context.Context, req *emptypb.Empty) (*shop_v1.ProductsGet_Response, error) {
	products, err := s.service.GetProducts(ctx)
	if err != nil {
		return &shop_v1.ProductsGet_Response{}, err
	}

	var response []*shop_v1.ShortProductInfo

	for _, product := range products {
		response = append(response, &shop_v1.ShortProductInfo{ProductId: product.ID.String(),
			Name:  product.Name,
			Price: float32(product.Price),
		})
	}
	return &shop_v1.ProductsGet_Response{Products: response}, nil
}
