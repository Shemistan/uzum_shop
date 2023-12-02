package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Phone       string      `json:"phone"`
	Role        string      `json:"role"`
	Address     string      `json:"address"`
	Coordinates Coordinates `json:"coordinates"`
}
