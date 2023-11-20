package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *shopSystemService) CancelOrderService(ctx context.Context, req *pb.CancelOrder_Request) (*empty.Empty, error) {
	resp, err := s.storage.CancelOrderDetailsStorage(ctx, int(req.OrderId))
	if err != nil {
		return nil, err
	}
	_, err = s.storage.CancelOrderStorage(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, v := range resp {
		stockCounts, err := s.storage.GetProductCountStorage(ctx, uint32(v.ProductId))
		if err != nil {
			return nil, err
		}
		countToSet := stockCounts + v.Count
		_, err = s.storage.CalculateProductCountStorage(ctx, v.ProductId, countToSet)
		if err != nil {
			return nil, err
		}
	}

	return nil, err
}
