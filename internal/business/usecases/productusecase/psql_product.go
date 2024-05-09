package productusecase

import (
	"DB_Project/internal/business/domains/productdomain"
	"database/sql"
	"fmt"
)

type PSQLProduct struct {
	DB *sql.DB
}

func NewPSQLProduct(db *sql.DB) productdomain.ProductDomain {
	return &PSQLProduct{
		DB: db,
	}
}

func (psql *PSQLProduct) GetAllProducts() ([]*productdomain.Product, error) {
	rows, err := psql.DB.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*productdomain.Product
	for rows.Next() {
		var p productdomain.Product
		err := rows.Scan(&p.ID, &p.CategoryName, &p.ManufacturerName, &p.Description, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (psql *PSQLProduct) GetProduct(id string) (*productdomain.Product, error) {
	row := psql.DB.QueryRow("SELECT * FROM product WHERE id = $1", id)

	var p productdomain.Product
	err := row.Scan(&p.ID, &p.CategoryName, &p.ManufacturerName, &p.Description, &p.Price, &p.Stock)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (psql *PSQLProduct) PostProduct(product *productdomain.Product) (int, error) {
	fmt.Printf("PostProduct called with product: %+v\n", product)

	var id int
	err := psql.DB.QueryRow(
		"INSERT INTO product (category_name, manufacturer_name, description, price, stock) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		product.CategoryName,
		product.ManufacturerName,
		product.Description,
		product.Price,
		product.Stock,
	).Scan(&id)

	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return 0, err
	}

	fmt.Printf("Product posted successfully with ID: %d\n", id)

	return id, nil
}
