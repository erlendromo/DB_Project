package showcasedomain

import (
	"time"
)

type ProductSales struct {
	ProductID   int
	Description string
	TotalSales  float64
}

type DiscountedProduct struct {
	ProductID           int
	Description         string
	DiscountPercentage  float64
	DiscountDescription string
	EndDate             string
}

type ProductInfo struct {
	Description string
	Quantity    int
}

type OrderDetail struct {
	OrderID         int
	PlacedAt        time.Time
	TotalAmount     float64
	Status          string
	Username        string
	FirstName       string
	LastName        string
	Email           string
	PhoneNumber     string
	PaymentCount    int
	PaymentStatuses string
	ProductsInfo    []ProductInfo
}

type TopCustomer struct {
	CustomerID     int
	Username       string
	NumberOfOrders int
	TotalSpent     float64
}
