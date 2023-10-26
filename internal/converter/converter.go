package converter

import (
	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"

	"github.com/Shemistan/uzum_shop/internal/models"
)

func ShortProductApiToModelService(req *models.Product) *pb.ShortProduct {
	return &pb.ShortProduct{
		Id:    req.ID,
		Name:  req.Name,
		Price: req.Price,
	}
}

func ProductApiToModelService(req *models.Product) *pb.Product {
	return &pb.Product{
		Id:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Count:       req.Count,
	}
}

func AddProductRequestApiToModelService(req *pb.Basket) *models.Basket {
	return &models.Basket{
		ID:        req.GetId(),
		UserID:    req.GetUserId(),
		ProductID: req.GetProductId(),
		Count:     req.GetCount(),
	}
}

func UpdateBasketRequestApiToModelService(req *pb.UpBasket) *models.UpdateBasketRequest {
	return &models.UpdateBasketRequest{
		ProductID: req.GetProductId(),
		Count:     req.GetCount(),
	}
}

func BasketModelServiceToApi(req *models.Basket) *pb.Basket {
	return &pb.Basket{
		UserId:    req.UserID,
		ProductId: req.ProductID,
		Count:     req.Count,
	}
}
