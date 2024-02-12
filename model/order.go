package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_itme"`
	OrderStatus string     `json:"order_status"`
	CreatedAt   time.Time  `json:"created_at"`
	ShippedAt   time.Time  `json:"shipped_at"`
	CompletedAt time.Time  `json:"completed_at"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"itme_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
