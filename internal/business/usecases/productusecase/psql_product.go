package productusecase

import (
	"DB_Project/internal/business/domains/productdomain"
	"database/sql"
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
