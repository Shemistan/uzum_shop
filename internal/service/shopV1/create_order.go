package shopV1

import (
	"context"
	"errors"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"time"
)

func (s *shopSystemService) CreateOrderService(ctx context.Context, req *pb.Order_Request) (*pb.Order_Response, error) {
	userId, err := s.GetUserIdFromLoginServ(ctx)
	if err != nil {
		return &pb.Order_Response{}, err
	}

	getItemsFromBasket, err := s.storage.GetItemsFromBasket(ctx, userId)
	if err != nil {
		return &pb.Order_Response{}, err
	}

	stockMap := make(map[int]int)

	for _, v := range getItemsFromBasket {
		stockCounts, err := s.storage.GetProductCountStorage(ctx, uint32(v.ProductId))
		if err != nil {
			return nil, err
		}

		if stockCounts < v.Count {
			return nil, errors.New("not enough stock")
		}
		stockMap[v.ProductId] = stockCounts
	}

	var address string = req.Address

	if req.Address == "" {
		address, err = s.storage.GetAddress(ctx, userId)
		if err != nil {
			return nil, err
		}

		if address == "" {
			return nil, errors.New("no address provided")
		}
	}

	orderModel := &models.Order{
		User_id: userId,
		//Products_id:          arr1,
		Address:              address,
		Coordinate_address_x: req.DropX,
		Coordinate_address_y: req.DropY,
		Coordinates_point_x:  req.TakeX,
		Coordinates_point_y:  req.TakeY,
		Create_at:            time.Now().Format("2006-01-02 15:04:05"),
		Delivery_status:      "Awaiting Shipment",
	}

	respOrderId, err := s.storage.CreateOrderStorage(ctx, orderModel)
	if err != nil {
		return &pb.Order_Response{}, err
	}

	err = s.storage.CreateOrderDetails(ctx, int(respOrderId), getItemsFromBasket)
	if err != nil {
		return &pb.Order_Response{}, err
	}

	for _, item := range getItemsFromBasket {

		deleteReq := &models.DeleteFomBasked{
			UserId:    userId,
			ProductId: uint32(item.ProductId),
		}

		err = s.storage.DeleteBasketStorage(ctx, deleteReq)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Order_Response{
		OrderId: respOrderId,
	}, nil
}
