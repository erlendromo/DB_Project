package customeraddressusecase

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"database/sql"
)

type PSQLAddress struct {
	DB *sql.DB
}

func NewPSQLAddress(db *sql.DB) customeraddressdomain.AddressDomain {
	return &PSQLAddress{
		DB: db,
	}
}

func (psql *PSQLAddress) CreateAddress(address *customeraddressdomain.CreateAddress) (*customeraddressdomain.Address, error) {
	return nil, nil
}

func (psql *PSQLAddress) GetAddressByID(id int) (*customeraddressdomain.Address, error) {
	return nil, nil
}

func (psql *PSQLAddress) UpdateAddress(id int, address *customeraddressdomain.CreateAddress) (*customeraddressdomain.Address, error) {
	return nil, nil
}

func (psql *PSQLAddress) SoftDeleteAddress(id int) error {
	return nil
}
