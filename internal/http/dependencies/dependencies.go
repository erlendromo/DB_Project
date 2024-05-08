package dependencies

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"DB_Project/internal/business/usecases/customeraddressusecase"
	"database/sql"
)

var Dependencies Deps

type Deps struct {
	CustomerDeps *CustomerDeps
}

func InitDeps(db *sql.DB) {
	Dependencies.CustomerDeps = newCustomerDeps(
		customeraddressusecase.NewPSQLCustomer(db),
		customeraddressusecase.NewPSQLAddress(db),
		customeraddressusecase.NewPSQLCustomerAddress(db),
	)
}

type CustomerDeps struct {
	CustomerDomain        customeraddressdomain.CustomerDomain
	AddressDomain         customeraddressdomain.AddressDomain
	CustomerAddressDomain customeraddressdomain.CustomerAddressDomain
}

func newCustomerDeps(cd customeraddressdomain.CustomerDomain, ad customeraddressdomain.AddressDomain, cad customeraddressdomain.CustomerAddressDomain) *CustomerDeps {
	return &CustomerDeps{
		CustomerDomain:        cd,
		AddressDomain:         ad,
		CustomerAddressDomain: cad,
	}
}
