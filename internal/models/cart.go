package models

import "github.com/google/uuid"

type Cart struct {
	Items map[uuid.UUID]int32
	Total float64
}
