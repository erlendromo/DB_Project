package customeraddressusecase

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"context"
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

func (psql *PSQLCustomer) CreateCustomer(ctx context.Context, customer *customeraddressdomain.CreateCustomer) (int, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var customerID int
	row := tx.QueryRow(`INSERT INTO customer (username, password, first_name, last_name, email, phone_number) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, customer.Username, customer.Password, customer.FirstName, customer.LastName, customer.Email, customer.PhoneNumber)
	if err := row.Scan(&customerID); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return customerID, nil
}

func (psql *PSQLCustomer) GetCustomerByID(ctx context.Context, id int) (*customeraddressdomain.DBCustomer, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	row := tx.QueryRow(`SELECT username, password, first_name, last_name, email, phone_number FROM customer WHERE id = $1`, id)

	var customer customeraddressdomain.DBCustomer
	if err := row.Scan(&customer.Username, &customer.Password, &customer.FirstName, &customer.LastName, &customer.Email, &customer.PhoneNumber); err != nil {
		return nil, err
	}

	return &customer, nil
}

func (psql *PSQLCustomer) GetCustomerByUsername(ctx context.Context, username string) (*customeraddressdomain.DBCustomer, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	row := tx.QueryRow(`SELECT id, password, first_name, last_name, email, phone_number, role FROM customer WHERE username = $1`, username)

	var customer customeraddressdomain.DBCustomer
	if err := row.Scan(&customer.ID, &customer.Password, &customer.FirstName, &customer.LastName, &customer.Email, &customer.PhoneNumber, &customer.Role); err != nil {
		return nil, err
	}

	return &customer, nil
}

func (psql *PSQLCustomer) UpdateCustomer(ctx context.Context, id int, customer *customeraddressdomain.CreateCustomer) error {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `UPDATE customer SET username = $1, password = $2, first_name = $3, last_name = $4, email = $5, phone_number = $6 WHERE id = $7`, customer.Username, customer.Password, customer.FirstName, customer.LastName, customer.Email, customer.PhoneNumber, id); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (psql *PSQLCustomer) SoftDeleteCustomer(ctx context.Context, id int) error {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "UPDATE customer SET deleted = true WHERE id = $1", id); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (psql *PSQLCustomer) GetAdminByUsername(ctx context.Context, username string) (bool, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}

	defer tx.Rollback()

	row := tx.QueryRow(`SELECT role FROM customer WHERE username = $1`, username)

	var role int
	if err := row.Scan(&role); err != nil {
		return false, err
	}

	if role != 1 {
		return false, nil
	}

	return true, nil
}
