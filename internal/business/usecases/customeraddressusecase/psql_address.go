package customeraddressusecase

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"context"
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

func (psql *PSQLAddress) CreateAddress(ctx context.Context, address *customeraddressdomain.CreateAddress) (int, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	createAddress, err := address.SetCity()
	if err != nil {
		return 0, err
	}

	if err := tx.QueryRowContext(ctx, "SELECT zip FROM zipcode WHERE zip = $1", createAddress.ZipCode).Scan(&createAddress.ZipCode); err != nil {
		if err == sql.ErrNoRows {

			if _, err := tx.ExecContext(ctx, "INSERT INTO zipcode (zip, city) VALUES ($1, $2)", createAddress.ZipCode, createAddress.City); err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	var addressID int
	row := tx.QueryRow(`INSERT INTO address (zipcode, street) VALUES ($1, $2) RETURNING id`, createAddress.ZipCode, createAddress.Street)
	if err := row.Scan(&addressID); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return addressID, nil
}

func (psql *PSQLAddress) GetAddressByID(ctx context.Context, id int) (*customeraddressdomain.Address, error) {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	row := tx.QueryRow(`SELECT a.zipcode, z.city, a.street FROM address a JOIN zipcode z ON a.zipcode = z.zip WHERE a.id = $1`, id)

	var address customeraddressdomain.Address
	if err := row.Scan(&address.ZipCode, &address.City, &address.Street); err != nil {
		return nil, err
	}

	return &address, nil
}

func (psql *PSQLAddress) SoftDeleteAddress(ctx context.Context, id int) error {
	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "UPDATE address SET deleted = true WHERE id = $1", id); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
