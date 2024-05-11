package shoppingorderusecase

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"DB_Project/internal/business/domains/shoppingorderdomain"
	"DB_Project/internal/constants"
	"context"
	"database/sql"
	"errors"
	"time"
)

type PSQLShoppingOrder struct {
	DB                  *sql.DB
	PSQLCustomerAddress customeraddressdomain.CustomerAddressDomain
}

func NewPSQLShoppingOrder(db *sql.DB, cad customeraddressdomain.CustomerAddressDomain) shoppingorderdomain.ShoppingOrderDomain {
	return &PSQLShoppingOrder{
		DB:                  db,
		PSQLCustomerAddress: cad,
	}
}

func (psql *PSQLShoppingOrder) CreateOrder(ctx context.Context, customerID int, items map[int]int) (*shoppingorderdomain.ShoppingOrderResponse, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	customerAddr, err := psql.PSQLCustomerAddress.GetCustomerAddressesByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	var shoppingOrderID int
	row := tx.QueryRowContext(ctx, "INSERT INTO shopping_order (customer_id, placed_at, total_amount) VALUES ($1, NOW(), 0.0) RETURNING id", customerID)
	if err := row.Scan(&shoppingOrderID); err != nil {
		return nil, err
	}

	responseItems := make([]shoppingorderdomain.ItemResponse, 0)
	totalAmount := 0.0
	subtotal := 0.0

	for productID, quantity := range items {
		row := tx.QueryRow("SELECT id, category_name, manufacturer_name, description, price, stock FROM product WHERE id = $1", productID)
		var product shoppingorderdomain.DBProduct
		if err := row.Scan(&product.ID, &product.CategoryName, &product.ManufacturerName, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}

		if product.Stock < quantity {
			return nil, errors.New("insufficient stock")
		}

		subtotal = product.Price * float64(quantity)

		if _, err := tx.Exec("INSERT INTO item (shopping_order_id, product_id, quantity, sub_total) VALUES ($1, $2, $3, $4)", shoppingOrderID, productID, quantity, subtotal); err != nil {
			return nil, err
		}

		if _, err := tx.Exec("UPDATE product SET stock = stock - $1 WHERE id = $2", quantity, productID); err != nil {
			return nil, err
		}

		responseItems = append(responseItems, shoppingorderdomain.ItemResponse{
			ProductID:          product.ID,
			ProductDescription: product.Description,
			Quantity:           quantity,
			Subtotal:           subtotal,
		})

		totalAmount += subtotal
	}

	if _, err := tx.Exec(`UPDATE shopping_order SET total_amount = $1, status = 'Created' WHERE id = $2`, totalAmount, shoppingOrderID); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	shoppingOrderResponse := &shoppingorderdomain.ShoppingOrderResponse{
		OrderID: shoppingOrderID,
		Customer: shoppingorderdomain.CustomerResponse{
			FirstName:   customerAddr.FirstName,
			LastName:    customerAddr.LastName,
			Email:       customerAddr.Email,
			PhoneNumber: customerAddr.PhoneNumber,
			Addresses:   customerAddr.Addresses,
		},
		PlacedAt:    time.Now().Format(constants.TIME_FORMAT),
		TotalAmount: totalAmount,
		Status:      "Created",
		Items:       responseItems,
	}

	return shoppingOrderResponse, nil
}

func (psql *PSQLShoppingOrder) GetOrderByID(ctx context.Context, customerID, shoppingOrderID int) string {
	// STEP 1: Query "customer_address" table with customerID
	// STEP 2: Query "item" table with shoppingOrderID
	// STEP 3: Query "product" table with productID from "item" table
	// STEP 4: Return the result

	return "GetOrderByID"
}

func (psql *PSQLShoppingOrder) GetOrders(ctx context.Context, customerID int) string {
	// STEP 1: Query "customer_address" table with customerID
	// STEP 2: Query "item" table with shoppingOrderID
	// STEP 3: Query "product" table with productID from "item" table
	// STEP 4: Populate a list of shopping orders
	// STEP 5: Return the result

	return "GetAllOrders"
}

func (psql *PSQLShoppingOrder) UpdateOrder(ctx context.Context, customerID, shoppingOrderID int, items map[int]int) string {
	// STEP 1: Update "shopping_order" table
	// STEP 2: Update "item" table
	// STEP 3: Return the result

	return "UpdateOrder"
}
