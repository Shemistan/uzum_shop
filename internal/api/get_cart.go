package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"uzum_shop/generated/protos/shop_v1"
)

func (s *ShopAPI) GetCart(ctx context.Context, req *emptypb.Empty) (*shop_v1.Cart_GetResponse, error) {
	cart, err := s.service.GetCart(ctx)
	if err != nil {
		return &shop_v1.Cart_GetResponse{}, err
	}

	var responseItems []*shop_v1.Product

	for _, item := range *cart {
		responseItems = append(responseItems, &shop_v1.Product{
			ProductId:   item.ID.String(),
			Name:        item.Name,
			Price:       float32(item.Price),
			Description: item.Description,
			Quantity:    uint32(item.Quantity)})
	}

	return &shop_v1.Cart_GetResponse{Products: responseItems}, nil
}
