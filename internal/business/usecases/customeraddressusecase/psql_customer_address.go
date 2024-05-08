package customeraddressusecase

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"database/sql"
)

type PSQLCustomerAddress struct {
	DB *sql.DB
}

func NewPSQLCustomerAddress(db *sql.DB) customeraddressdomain.CustomerAddressDomain {
	return &PSQLCustomerAddress{
		DB: db,
	}
}

func (psql *PSQLCustomerAddress) CreateCustomerAddress(customerID int, addressID int, primaryAddress bool) (*customeraddressdomain.DBCustomerAddress, error) {
	return nil, nil
}

func (psql *PSQLCustomerAddress) GetCustomerAddressByCustomerID(customerID int) (*customeraddressdomain.DBCustomerAddress, error) {
	return nil, nil
}

func (psql *PSQLCustomerAddress) GetCustomerPrimaryAddressByCustomerID(customerID int) (*customeraddressdomain.DBCustomerAddress, error) {
	return nil, nil
}

func (psql *PSQLCustomerAddress) UpdatePrimaryAddress(customerID int, addressID int) (*customeraddressdomain.DBCustomerAddress, error) {
	return nil, nil
}

// ADMINS ONLY
func (psql *PSQLCustomerAddress) AllCustomersAddress() ([]customeraddressdomain.DBCustomerAddress, error) {
	return nil, nil
}
