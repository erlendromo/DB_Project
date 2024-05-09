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

	return id, nil
}

func (psql *PSQLProduct) SearchProductFullText(description string) ([]*productdomain.Product, error) {
	query := `
		SELECT * FROM product 
		WHERE to_tsvector('english', description) @@ plainto_tsquery('english', $1)
	`
	rows, err := psql.DB.Query(query, description)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var products []*productdomain.Product
	for rows.Next() {
		var product productdomain.Product
		err = rows.Scan(&product.ID, &product.CategoryName, &product.ManufacturerName, &product.Description, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (psql *PSQLProduct) PatchProduct(id string, product *productdomain.Product) (*productdomain.Product, error) {
	query := `
		UPDATE product 
		SET category_name = COALESCE($1, category_name), 
			manufacturer_name = COALESCE($2, manufacturer_name), 
			description = COALESCE($3, description), 
			price = COALESCE($4, price), 
			stock = COALESCE($5, stock)
		WHERE id = $6
		RETURNING *
	`
	row := psql.DB.QueryRow(query, product.CategoryName, product.ManufacturerName, product.Description, product.Price, product.Stock, id)

	newProduct := &productdomain.Product{}
	err := row.Scan(&newProduct.ID, &newProduct.CategoryName, &newProduct.ManufacturerName, &newProduct.Description, &newProduct.Price, &newProduct.Stock)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return nil, err
	}

	return newProduct, nil
}

func (psql *PSQLProduct) DeleteProduct(id string) error {
	query := "DELETE FROM product WHERE id = $1"
	_, err := psql.DB.Exec(query, id)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return err
	}

	return nil
}
