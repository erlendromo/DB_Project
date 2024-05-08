package customeraddressdomain

import "errors"

type AddressDomain interface {
	CreateAddress(address *CreateAddress) (*Address, error)
	GetAddressByID(id int) (*Address, error)
	UpdateAddress(id int, address *CreateAddress) (*Address, error)
	SoftDeleteAddress(id int) error
}

type DBZip struct {
	Zip  string `db:"zip"`
	City string `db:"city"`
}

type DBAddress struct {
	ID      int    `db:"id"`
	ZipCode DBZip  `db:"zip_code"`
	Street  string `db:"street"`
	Deleted bool   `db:"deleted"`
}

type Address struct {
	ZipCode int    `json:"zip_code"`
	Street  string `json:"street"`
}

type CreateZip struct {
	Zip  string `json:"zip"`
	City string `json:"city"`
}

func (c *CreateZip) Validate() error {
	if len(c.Zip) != 4 {
		return errors.New("zip must be 4 characters")
	}

	if len(c.City) < 4 {
		return errors.New("city must be at least 4 characters")
	}

	return nil
}

type CreateAddress struct {
	ZipCode string `json:"zip_code"`
	Street  string `json:"street"`
}

func (c *CreateAddress) Validate() error {
	if len(c.ZipCode) < 4 {
		return errors.New("zip code must be at least 4 characters")
	}

	if len(c.Street) < 4 {
		return errors.New("street must be at least 4 characters")
	}

	// TODO add add of city for specific zipcodes? (use map with zipcodes as keys and cities as values)

	return nil
}
