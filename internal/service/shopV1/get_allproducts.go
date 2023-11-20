package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (s *shopSystemService) GetAllProductsService(ctx context.Context) (*pb.GetAllProducts_Response, error) {
	response, err := s.storage.GetAllProductsStorage(ctx)
	if err != nil {
		return &pb.GetAllProducts_Response{}, err
	}

	var all []*pb.ProductShort
	for _, p := range response {
		all = append(all, &pb.ProductShort{
			Name:  p.Name,
			Price: p.Price,
		})
	}
	return &pb.GetAllProducts_Response{
		AllProducts: all,
	}, nil
}
