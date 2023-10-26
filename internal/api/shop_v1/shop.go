package shop_v1

import (
	"github.com/Shemistan/uzum_shop/internal/service"
	pb "github.com/Shemistan/uzum_shop/pkg/shop_v1"
)

type Shop struct {
	pb.UnimplementedShopV1Server

	ShopService service.IService
}
