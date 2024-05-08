package customeraddressdomain

type DBCustomerAddress struct {
	CustomerID     int  `db:"customer_id"`
	AddressID      int  `db:"address_id"`
	PrimaryAddress bool `db:"primary_address"`
}

type CustomerAddressDomain interface {
	CreateCustomerAddress(customerID int, addressID int, primaryAddress bool) (*DBCustomerAddress, error)
	GetCustomerAddressByCustomerID(customerID int) (*DBCustomerAddress, error)
	GetCustomerPrimaryAddressByCustomerID(customerID int) (*DBCustomerAddress, error)
	UpdatePrimaryAddress(customerID int, addressID int) (*DBCustomerAddress, error)

	// ADMINS ONLY
	AllCustomersAddress() ([]DBCustomerAddress, error)
}
