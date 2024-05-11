package shoppingorderdomain

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"context"
)

type ShoppingOrderDomain interface {
	CreateOrder(ctx context.Context, customerID int, items map[int]int) (*ShoppingOrderResponse, error)
	GetOrderByID(ctx context.Context, customerID, shoppingOrderID int) string
	GetOrders(ctx context.Context, customerID int) string
	UpdateOrder(ctx context.Context, customerID, shoppingOrderID int, items map[int]int) string
}

// Database

type DBShoppingOrder struct {
	ID          int     `db:"id"`
	CustomerID  int     `db:"customer_id"`
	PlacedAt    string  `db:"placed_at"`
	TotalAmount float64 `db:"total_amount"`
	Status      string  `db:"status"`
}

type DBItem struct {
	ID              int     `db:"id"`
	ShoppingOrderID int     `db:"shopping_order_id"`
	ProductID       int     `db:"product_id"`
	Quantity        int     `db:"quantity"`
	Subtotal        float64 `db:"subtotal"`
}

type DBProduct struct {
	ID               int     `db:"id"`
	CategoryName     string  `db:"category_name"`
	ManufacturerName string  `db:"manufacturer_name"`
	Description      string  `db:"description"`
	Price            float64 `db:"price"`
	Stock            int     `db:"stock"`
}

// Requests

type CreateItemRequest struct {
	ShoppingOrderID int     `json:"shopping_order_id"`
	ProductID       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	Subtotal        float64 `json:"subtotal"`
}

// Responses

type ShoppingOrderResponse struct {
	OrderID     int              `json:"order_id"`
	Customer    CustomerResponse `json:"customer"`
	PlacedAt    string           `json:"placed_at"`
	TotalAmount float64          `json:"total_amount"`
	Status      string           `json:"status"`
	Items       []ItemResponse   `json:"items"`
}

type CustomerResponse struct {
	FirstName   string                          `json:"first_name"`
	LastName    string                          `json:"last_name"`
	Email       string                          `json:"email"`
	PhoneNumber string                          `json:"phone_number"`
	Addresses   []customeraddressdomain.Address `json:"addresses"`
}

type ItemResponse struct {
	ProductID          int     `json:"product_id"`
	ProductDescription string  `json:"product_description"`
	Quantity           int     `json:"quantity"`
	Subtotal           float64 `json:"subtotal"`
}
