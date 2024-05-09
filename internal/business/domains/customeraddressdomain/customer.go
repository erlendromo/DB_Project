package customeraddressdomain

import (
	"context"
	"errors"
	"strings"
)

type CustomerDomain interface {
	CreateCustomer(ctx context.Context, customer *CreateCustomer) (*Customer, error)
	GetCustomerByUsername(ctx context.Context, username string) (*Customer, error)
	UpdateCustomer(ctx context.Context, username string, customer *CreateCustomer) (*Customer, error)
	SoftDeleteCustomer(ctx context.Context, username string) error
}

type DBCustomer struct {
	ID          int    `db:"id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	Role        int16  `db:"role"`
	Deleted     bool   `db:"deleted"`
}

type Customer struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CreateCustomer struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (c *CreateCustomer) Validate() error {
	if len(c.Username) < 4 {
		return errors.New("username must be at least 4 characters")
	} else if len(c.Username) > 50 {
		return errors.New("username must be at most 50 characters")
	}

	if len(c.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	} else if len(c.Password) > 50 {
		return errors.New("password must be at most 50 characters")
	}

	if len(c.FirstName) < 2 {
		return errors.New("first name must be at least 2 characters")
	} else if len(c.FirstName) > 50 {
		return errors.New("first name must be at most 50 characters")
	}

	if len(c.LastName) < 2 {
		return errors.New("last name must be at least 2 characters")
	} else if len(c.LastName) > 50 {
		return errors.New("last name must be at most 50 characters")
	}

	if !strings.Contains(c.Email, "@") {
		return errors.New("invalid email")
	} else if len(c.Email) < 8 {
		return errors.New("email must be at least 8 characters")
	} else if len(c.Email) > 50 {
		return errors.New("email must be at most 50 characters")
	}

	if len(c.PhoneNumber) < 8 {
		return errors.New("phone number must be at least 8 characters")
	} else if len(c.PhoneNumber) > 50 {
		return errors.New("phone number must be at most 15 characters")
	}

	return nil
}
