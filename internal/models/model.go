package models

import "github.com/golang-jwt/jwt"

type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Count       uint32  `json:"count,omitempty"`
}

type AddProductToBasketModel struct {
	UserId    int
	ProductId uint32
	Count     uint32
}

type GetFromBasket struct {
	ProductId int
	Count     int
}

type DeleteFomBasked struct {
	UserId    int
	ProductId uint32
}

type Order struct {
	ID                   int
	User_id              int
	Products_id          []int
	Address              string
	Coordinate_address_x float64
	Coordinate_address_y float64
	Coordinates_point_x  float64
	Coordinates_point_y  float64
	Create_at            string
	//Start_at             string
	//Delivery_at          string
	//Courier_id           int
	Delivery_status string
}

type CustomClaims struct {
	jwt.StandardClaims
	Login string `json:"username"`
	Role  string `json:"role"`
}
