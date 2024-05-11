package customeraddressdomain

import (
	"context"
	"errors"
	"strconv"
)

type DBAddress struct {
	ID      int    `db:"id"`
	ZipCode string `db:"zipcode"`
	Street  string `db:"street"`
	Deleted bool   `db:"deleted"`
}

type Address struct {
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type CreateAddress struct {
	ZipCode string `json:"zip_code"`
	Street  string `json:"street"`
}

func (c *CreateAddress) SetCity() (*Address, error) {
	a := &Address{
		ZipCode: c.ZipCode,
		Street:  c.Street,
	}

	zipInt, err := strconv.Atoi(c.ZipCode)
	if err != nil {
		return nil, err
	}

	if zipInt < 10 {
		return nil, errors.New("invalid zip code")
	} else if zipInt < 2000 {
		a.City = "Oslo"
	} else if zipInt < 3000 {
		a.City = "Lillestrøm"
	} else if zipInt < 4000 {
		a.City = "Drammen"
	} else if zipInt < 5000 {
		a.City = "Stavanger"
	} else if zipInt < 6000 {
		a.City = "Bergen"
	} else if zipInt < 7000 {
		a.City = "Ålesund"
	} else if zipInt < 8000 {
		a.City = "Trondheim"
	} else if zipInt < 9000 {
		a.City = "Bodø"
	} else if zipInt < 10000 {
		a.City = "Tromsø"
	} else {
		return nil, errors.New("invalid zip code")
	}

	return a, nil
}

type AddressDomain interface {
	CreateAddress(ctx context.Context, address *CreateAddress) (int, error)
	GetAddressByID(ctx context.Context, id int) (*Address, error)
	SoftDeleteAddress(ctx context.Context, id int) error
}
