package response

import (
	"time"
)

type ProductCreateResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
	CreatedAt  time.Time
}

type ProductGetResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProductUpdateResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
