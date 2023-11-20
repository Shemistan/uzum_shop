package shopV1

import (
	shop_system "github.com/Shemistan/uzum_shop/internal/service/shopV1"
	"github.com/Shemistan/uzum_shop/pkg/shopV1"
)

type Shop struct {
	shopV1.UnimplementedShopServer
	ShopService shop_system.IShopSystemService
}
