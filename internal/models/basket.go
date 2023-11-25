package models

type Basket struct {
	ID        int64  `json:"id"`
	UserID    string `json:"user_id"`
	ProductID int64  `json:"product_id"`
	Count     int32  `json:"count"`
}

type UpdateBasketRequest struct {
	ProductID int64 `json:"product_id"`
	Count     int32 `json:"count"`
}
