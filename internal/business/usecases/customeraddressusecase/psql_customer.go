package customeraddressusecase

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"database/sql"
)

type PSQLCustomer struct {
	DB *sql.DB
}

func NewPSQLCustomer(db *sql.DB) customeraddressdomain.CustomerDomain {
	return &PSQLCustomer{
		DB: db,
	}
}

func (psql *PSQLCustomer) CreateCustomer(customer *customeraddressdomain.CreateCustomer) (*customeraddressdomain.Customer, error) {
	return nil, nil
}

func (psql *PSQLCustomer) GetCustomerByUsername(username string) (*customeraddressdomain.Customer, error) {
	return nil, nil
}

func (psql *PSQLCustomer) UpdateCustomer(username string, customer *customeraddressdomain.CreateCustomer) (*customeraddressdomain.Customer, error) {
	return nil, nil
}

func (psql *PSQLCustomer) SoftDeleteCustomer(username string) error {
	return nil
}
