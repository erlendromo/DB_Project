package showcaseusecase

import (
	"DB_Project/internal/business/domains/showcasedomain"
	"database/sql"
	"encoding/json"
	"fmt"
)

type PSQLShowcase struct {
	DB *sql.DB
}

func NewPSQLShowcase(db *sql.DB) showcasedomain.ShowcaseDomain {
	return &PSQLShowcase{
		DB: db,
	}
}

func (psql *PSQLShowcase) CalculateTotalSalesPerProduct() ([]*showcasedomain.ProductSales, error) {
	stmt, err := psql.DB.Prepare(`
        SELECT p.id, p.description, SUM(i.sub_total) AS total_sales
        FROM product p
        JOIN item i ON p.id = i.product_id
        GROUP BY p.id, p.description;
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []*showcasedomain.ProductSales
	for rows.Next() {
		var s showcasedomain.ProductSales
		if err := rows.Scan(&s.ProductID, &s.Description, &s.TotalSales); err != nil {
			return nil, err
		}
		sales = append(sales, &s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sales, nil
}

func (psql *PSQLShowcase) ListCurrentDiscountedProducts() ([]*showcasedomain.DiscountedProduct, error) {
	stmt, err := psql.DB.Prepare(`
        SELECT p.id, p.description, d.percentage, d.description AS discount_description, d.end_at
        FROM product p
        JOIN product_discount pd ON p.id = pd.product_id
        JOIN discount d ON pd.discount_id = d.id
        WHERE CURRENT_TIMESTAMP BETWEEN d.start_at AND d.end_at;
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discountedProducts []*showcasedomain.DiscountedProduct
	for rows.Next() {
		var dp showcasedomain.DiscountedProduct
		if err := rows.Scan(&dp.ProductID, &dp.Description, &dp.DiscountPercentage, &dp.DiscountDescription, &dp.EndDate); err != nil {
			return nil, err
		}
		discountedProducts = append(discountedProducts, &dp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return discountedProducts, nil
}

func (psql *PSQLShowcase) FetchOrderWithDetails(orderID string) (*showcasedomain.OrderDetail, error) {
	stmt, err := psql.DB.Prepare(`
        SELECT 
            o.id, 
            o.placed_at, 
            o.total_amount, 
            o.status AS order_status, 
            c.username, 
            c.first_name, 
            c.last_name, 
            c.email, 
            c.phone_number,
            COUNT(p.id) AS payment_count,
            STRING_AGG(DISTINCT p.status, ', ') AS payment_statuses,
            json_agg(json_build_object('description', pr.description, 'quantity', i.quantity)) AS products_info
        FROM shopping_order o
        JOIN customer c ON o.customer_id = c.id
        LEFT JOIN payment p ON o.id = p.shopping_order_id
        LEFT JOIN item i ON o.id = i.shopping_order_id
        LEFT JOIN product pr ON i.product_id = pr.id
        WHERE o.id = $1
        GROUP BY o.id, c.id;
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(orderID)

	var od showcasedomain.OrderDetail
	var productsInfoJSON []byte
	var paymentStatuses string

	if err := row.Scan(
		&od.OrderID,
		&od.PlacedAt,
		&od.TotalAmount,
		&od.Status,
		&od.Username,
		&od.FirstName,
		&od.LastName,
		&od.Email,
		&od.PhoneNumber,
		&od.PaymentCount,
		&paymentStatuses,
		&productsInfoJSON,
	); err != nil {
		return nil, err
	}

	// Parse the JSON array of product info
	if err := json.Unmarshal(productsInfoJSON, &od.ProductsInfo); err != nil {
		return nil, fmt.Errorf("error parsing product info: %w", err)
	}

	od.PaymentStatuses = paymentStatuses

	return &od, nil
}

func (psql *PSQLShowcase) IdentifyTopCustomers(limit string) ([]*showcasedomain.TopCustomer, error) {
	stmt, err := psql.DB.Prepare(`
        SELECT c.id, c.username, COUNT(o.id) AS number_of_orders, SUM(o.total_amount) AS total_spent
        FROM customer c
        JOIN shopping_order o ON c.id = o.customer_id
        GROUP BY c.id
        ORDER BY SUM(o.total_amount) DESC
        LIMIT $1;
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topCustomers []*showcasedomain.TopCustomer
	for rows.Next() {
		var tc showcasedomain.TopCustomer
		if err := rows.Scan(&tc.CustomerID, &tc.Username, &tc.NumberOfOrders, &tc.TotalSpent); err != nil {
			return nil, err
		}
		topCustomers = append(topCustomers, &tc)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return topCustomers, nil
}
