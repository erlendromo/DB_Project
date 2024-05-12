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

func (psql *PSQLShoppingOrder) GetOrderByID(ctx context.Context, customerID, shoppingOrderID int) (*shoppingorderdomain.ShoppingOrderResponse, error) {
	// STEP 1: Query "customer_address" table with customerID
	// STEP 2: Query "item" table with shoppingOrderID
	// STEP 3: Query "product" table with productID from "item" table
	// STEP 4: Return the result

	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	customerAddr, err := psql.PSQLCustomerAddress.GetCustomerAddressesByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(`SELECT id, placed_at, total_amount, status FROM shopping_order WHERE id = $1`, shoppingOrderID)
	var shoppingOrder shoppingorderdomain.DBShoppingOrder
	if err := row.Scan(&shoppingOrder.ID, &shoppingOrder.PlacedAt, &shoppingOrder.TotalAmount, &shoppingOrder.Status); err != nil {
		return nil, err
	}

	rows, err := tx.Query(`SELECT id, shopping_order_id, product_id, quantity, sub_total FROM item WHERE shopping_order_id = $1`, shoppingOrderID)
	if err != nil {
		return nil, err
	}

	responseItems := make([]shoppingorderdomain.ItemResponse, 0)
	for rows.Next() {
		var item shoppingorderdomain.DBItem
		if err := rows.Scan(&item.ID, &item.ShoppingOrderID, &item.ProductID, &item.Quantity, &item.Subtotal); err != nil {
			return nil, err
		}

		tx2, err := psql.DB.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		defer tx.Rollback()

		row := tx2.QueryRow("SELECT id, category_name, manufacturer_name, description, price, stock FROM product WHERE id = $1", &item.ProductID)
		var product shoppingorderdomain.DBProduct
		if err := row.Scan(&product.ID, &product.CategoryName, &product.ManufacturerName, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}

		responseItems = append(responseItems, shoppingorderdomain.ItemResponse{
			ProductID:          item.ProductID,
			ProductDescription: product.Description,
			Quantity:           item.Quantity,
			Subtotal:           item.Subtotal,
		})

		if err := tx2.Commit(); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	shoppingOrderResponse := &shoppingorderdomain.ShoppingOrderResponse{
		OrderID: shoppingOrder.ID,
		Customer: shoppingorderdomain.CustomerResponse{
			FirstName:   customerAddr.FirstName,
			LastName:    customerAddr.LastName,
			Email:       customerAddr.Email,
			PhoneNumber: customerAddr.PhoneNumber,
			Addresses:   customerAddr.Addresses,
		},
		PlacedAt:    shoppingOrder.PlacedAt,
		TotalAmount: shoppingOrder.TotalAmount,
		Status:      shoppingOrder.Status,
		Items:       responseItems,
	}

	return shoppingOrderResponse, nil

}

func (psql *PSQLShoppingOrder) GetOrders(ctx context.Context, customerID int) ([]*shoppingorderdomain.ShoppingOrderResponse, error) {
	// STEP 1: Query "customer_address" table with customerID
	// STEP 2: Query "item" table with shoppingOrderID
	// STEP 3: Query "product" table with productID from "item" table
	// STEP 4: Populate a list of shopping orders
	// STEP 5: Return the result

	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT id FROM shopping_order WHERE customer_id = $1", customerID)
	if err != nil {
		return nil, err
	}

	shoppingOrderResponses := make([]*shoppingorderdomain.ShoppingOrderResponse, 0)
	for rows.Next() {
		var shoppingOrderID int
		if err := rows.Scan(&shoppingOrderID); err != nil {
			return nil, err
		}

		shoppingOrderResponse, err := psql.GetOrderByID(ctx, customerID, shoppingOrderID)
		if err != nil {
			return nil, err
		}

		shoppingOrderResponses = append(shoppingOrderResponses, shoppingOrderResponse)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return shoppingOrderResponses, nil
}
