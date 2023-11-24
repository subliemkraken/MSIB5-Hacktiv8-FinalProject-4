package response

import (
	"time"
)

type CategoryCreateResponse struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type CategoryGetResponse struct {
	ID                int        `json:"id"`
	Type              string     `json:"type"`
	SoldProductAmount int        `json:"sold_product_amount"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	Products          []Products `json:"products"`
}

type Products struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryUpdateResponse struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}
