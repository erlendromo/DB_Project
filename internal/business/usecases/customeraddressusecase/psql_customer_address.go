package customeraddressusecase

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"context"
	"database/sql"
)

type PSQLCustomerAddress struct {
	DB           *sql.DB
	PSQLCustomer customeraddressdomain.CustomerDomain
	PSQLAddress  customeraddressdomain.AddressDomain
}

func NewPSQLCustomerAddress(db *sql.DB, c customeraddressdomain.CustomerDomain, a customeraddressdomain.AddressDomain) customeraddressdomain.CustomerAddressDomain {
	return &PSQLCustomerAddress{
		DB:           db,
		PSQLCustomer: c,
		PSQLAddress:  a,
	}
}

func (psql *PSQLCustomerAddress) CreateCustomerAddress(ctx context.Context, custAddrReq *customeraddressdomain.CreateCustomerAddressRequest) (*customeraddressdomain.DBCustomerAddress, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	customerid, err := psql.PSQLCustomer.CreateCustomer(ctx,
		&customeraddressdomain.CreateCustomer{
			Username:    custAddrReq.Username,
			Password:    custAddrReq.Password,
			FirstName:   custAddrReq.FirstName,
			LastName:    custAddrReq.LastName,
			Email:       custAddrReq.Email,
			PhoneNumber: custAddrReq.PhoneNumber,
		},
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	addressid, err := psql.PSQLAddress.CreateAddress(ctx,
		&customeraddressdomain.CreateAddress{
			ZipCode: custAddrReq.ZipCode,
			Street:  custAddrReq.Street,
		},
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	customerAddressReq := &customeraddressdomain.DBCustomerAddress{
		CustomerID:     customerid,
		AddressID:      addressid,
		PrimaryAddress: custAddrReq.PrimaryAddress,
	}

	if _, err = psql.DB.ExecContext(ctx, `INSERT INTO customer_address (customer_id, address_id, primary_address) VALUES ($1, $2, $3)`, customerAddressReq.CustomerID, customerAddressReq.AddressID, customerAddressReq.PrimaryAddress); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return customerAddressReq, nil
}

func (psql *PSQLCustomerAddress) GetCustomerAddressesByCustomerID(ctx context.Context, customerID int) (*customeraddressdomain.CustomerAddresses, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `SELECT customer_id, address_id, primary_address FROM customer_address WHERE customer_id = $1`, customerID)
	if err != nil {
		return nil, err
	}

	var dbCustomerAddresses []customeraddressdomain.DBCustomerAddress
	for rows.Next() {
		var dbCustomerAddress customeraddressdomain.DBCustomerAddress

		if err = rows.Scan(&dbCustomerAddress.CustomerID, &dbCustomerAddress.AddressID, &dbCustomerAddress.PrimaryAddress); err != nil {
			return nil, err
		}

		dbCustomerAddresses = append(dbCustomerAddresses, dbCustomerAddress)
	}

	customer, err := psql.PSQLCustomer.GetCustomerByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	customerAddresses := customeraddressdomain.CustomerAddresses{
		Username:    customer.Username,
		Password:    customer.Password,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Addresses:   make([]customeraddressdomain.Address, 0),
	}

	for _, dbCustomerAddress := range dbCustomerAddresses {
		address, err := psql.PSQLAddress.GetAddressByID(ctx, dbCustomerAddress.AddressID)
		if err != nil {
			return nil, err
		}

		customerAddresses.Addresses = append(customerAddresses.Addresses, *address)
	}

	return &customerAddresses, nil
}

func (psql *PSQLCustomerAddress) GetCustomerPrimaryAddressByCustomerID(ctx context.Context, customerID int) (*customeraddressdomain.DBCustomerAddress, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var customerAddress customeraddressdomain.DBCustomerAddress
	row := tx.QueryRowContext(ctx, `SELECT customer_id, address_id, primary_address FROM customer_address WHERE customer_id = $1 AND primary_address = TRUE`, customerID)
	if err := row.Scan(&customerAddress.CustomerID, &customerAddress.AddressID, &customerAddress.PrimaryAddress); err != nil {
		return nil, err
	}

	return &customerAddress, nil
}

func (psql *PSQLCustomerAddress) UpdatePrimaryAddress(ctx context.Context, customerID int, addressID int, primary bool) (*customeraddressdomain.DBCustomerAddress, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var customerAddress customeraddressdomain.DBCustomerAddress
	row := tx.QueryRowContext(ctx, `UPDATE customer_address SET primary_address = $1 WHERE customer_id = $2 AND address_id = $3`, primary, customerID, addressID)
	if err := row.Scan(&customerAddress.CustomerID, &customerAddress.AddressID, &customerAddress.PrimaryAddress); err != nil {
		return nil, err
	}

	return &customerAddress, nil
}

// ADMINS ONLY
func (psql *PSQLCustomerAddress) AllCustomersAddresses(ctx context.Context) ([]customeraddressdomain.DBCustomerAddress, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `SELECT customer_id, address_id, primary_address FROM customer_address`)
	if err != nil {
		return nil, err
	}

	customerAddresses := make([]customeraddressdomain.DBCustomerAddress, 0)
	for rows.Next() {
		var customerAddress customeraddressdomain.DBCustomerAddress
		if err = rows.Scan(&customerAddress.CustomerID, &customerAddress.AddressID, &customerAddress.PrimaryAddress); err != nil {
			return nil, err
		}
		customerAddresses = append(customerAddresses, customerAddress)
	}

	return customerAddresses, nil
}
