package customeraddressdomain

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

type DBCustomerAddress struct {
	CustomerID     int  `db:"customer_id"`
	AddressID      int  `db:"address_id"`
	PrimaryAddress bool `db:"primary_address"`
}

type CustomerAddresses struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Addresses   []Address `json:"addresses"`
}

type CreateCustomerAddressRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	ZipCode        string `json:"zip_code"`
	Street         string `json:"street"`
	PrimaryAddress bool   `json:"primary_address"`
}

func FormatErrTooShort(key string) error {
	return errors.New(key + " is too short")
}

func FormatErrTooLong(key string) error {
	return errors.New(key + " is too long")
}

func CheckInput(input string, key string, min int, max int) error {
	if len(input) < min {
		return FormatErrTooShort(key)
	} else if len(input) > max {
		return FormatErrTooLong(key)
	}

	return nil
}

func (c *CreateCustomerAddressRequest) Validate() map[string]error {
	errs := make(map[string]error)

	if err := CheckInput(c.Username, "username", 4, 50); err != nil {
		errs["username"] = err
	}

	if err := CheckInput(c.Password, "password", 8, 50); err != nil {
		errs["password"] = err
	}

	if err := CheckInput(c.FirstName, "first_name", 2, 50); err != nil {
		errs["first_name"] = err
	}

	if err := CheckInput(c.LastName, "last_name", 2, 50); err != nil {
		errs["last_name"] = err
	}

	if err := CheckInput(c.Email, "email", 5, 50); err != nil {
		errs["email"] = err
	} else if !strings.Contains(c.Email, "@") {
		errs["email"] = errors.New("email is invalid")
	}

	if err := CheckInput(c.PhoneNumber, "phone_number", 5, 50); err != nil {
		errs["phone_number"] = err
	}

	if len(c.ZipCode) != 4 {
		errs["zip_code"] = errors.New("zip_code is invalid")
	} else if _, err := strconv.Atoi(c.ZipCode); err != nil {
		errs["zip_code"] = errors.New("zip_code must be a number")
	}

	if err := CheckInput(c.Street, "street", 5, 50); err != nil {
		errs["street"] = err
	}

	return errs
}

type CustomerAddressDomain interface {
	CreateCustomerAddress(ctx context.Context, custAddrReq *CreateCustomerAddressRequest) (*DBCustomerAddress, error)
	GetCustomerAddressesByCustomerID(ctx context.Context, customerID int) (*CustomerAddresses, error)
	UpdatePrimaryAddress(ctx context.Context, customerID int, addressID int, primary bool) (*DBCustomerAddress, error)

	// ADMINS ONLY
	AllCustomersAddresses(ctx context.Context) ([]DBCustomerAddress, error)
}
