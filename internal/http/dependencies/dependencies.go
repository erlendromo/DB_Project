package dependencies

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"DB_Project/internal/business/usecases/customeraddressusecase"
	"database/sql"
)

// All dependencies
var Dependencies Deps

type Deps struct {
	CustomerAddressDeps
}

func InitDeps(db *sql.DB) {
	c := customeraddressusecase.NewPSQLCustomer(db)
	a := customeraddressusecase.NewPSQLAddress(db)
	Dependencies.CustomerAddressDeps = CustomerAddressDeps{
		PSQLCustomerAddress: customeraddressusecase.NewPSQLCustomerAddress(db, c, a),
		PSQLCustomer:        c,
		PSQLAddress:         a,
	}
}

type CustomerAddressDeps struct {
	PSQLCustomerAddress customeraddressdomain.CustomerAddressDomain
	PSQLCustomer        customeraddressdomain.CustomerDomain
	PSQLAddress         customeraddressdomain.AddressDomain
}
