package userusecase

import (
	userdomain "DB_Project/internal/business/domains/user_domain"
	"context"
	"database/sql"
)

type PSQL_User struct {
	DB *sql.DB
}

func NewPSQLUser(db *sql.DB) userdomain.UserDomain {
	return &PSQL_User{
		DB: db,
	}
}

func (psqlu *PSQL_User) Insert(ctx context.Context, user *userdomain.User) (err error) {
	return
}

func (psqlu *PSQL_User) SelectByID(ctx context.Context, id int) (user userdomain.User, err error) {
	tx, err := psqlu.DB.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	row := tx.QueryRowContext(ctx,
		`SELECT u.*, a.* FROM user u 
		JOIN user_address ua ON u.user_id = ua.user_id
		JOIN address a ON ua.address_id = a.address_id
		WHERE u.user_id = $1;`, id,
	)

	if err = row.Scan(user.ID, user.Username, user.Password, user.FirstName,
		user.LastName, user.Email, user.PhoneNumber, user.Role); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
	}

	return
}

func (psqlu *PSQL_User) SelectAll(ctx context.Context) (users []userdomain.User, err error) {
	return
}
