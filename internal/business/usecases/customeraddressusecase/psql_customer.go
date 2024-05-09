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

func (psql *PSQLCustomer) CreateCustomer(ctx context.Context, customer *customeraddressdomain.CreateCustomer) (*customeraddressdomain.Customer, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if _, err = tx.ExecContext(ctx, `INSERT INTO "customer" (username, password, first_name, last_name, email, phone_number) VALUES ($1, $2, $3, $4, $5, $6)`,
		customer.Username, customer.Password, customer.FirstName, customer.LastName, customer.Email, customer.PhoneNumber); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (psql *PSQLCustomer) GetCustomerByUsername(ctx context.Context, username string) (*customeraddressdomain.Customer, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if _, err = tx.ExecContext(ctx, `SELECT first_name, last_name, email, phone_number FROM "customer" WHERE username = $1 AND deleted = FALSE`, username); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (psql *PSQLCustomer) UpdateCustomer(ctx context.Context, username string, customer *customeraddressdomain.CreateCustomer) (*customeraddressdomain.Customer, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	if _, err = tx.ExecContext(ctx, `UPDATE "customer" SET first_name = $1, last_name = $2, email = $3, phone_number = $4 WHERE username = $5`,
		customer.FirstName, customer.LastName, customer.Email, customer.PhoneNumber, username); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (psql *PSQLCustomer) SoftDeleteCustomer(ctx context.Context, username string) error {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, `UPDATE "customer" SET deleted = true WHERE username = $1`, username); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
