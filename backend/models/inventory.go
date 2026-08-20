package models

import "time"

// InventoryItem represents store stock products
type InventoryItem struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	ItemName      string    `json:"item_name"`
	Quantity      int       `json:"quantity"`
	UnitPrice     float64   `json:"unit_price"`
	LowStockLimit int       `json:"low_stock_limit"`
	CreatedAt     time.Time `json:"created_at"`
}