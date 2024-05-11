package dependencies

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"DB_Project/internal/business/domains/productdomain"
	"DB_Project/internal/business/domains/showcasedomain"
	"DB_Project/internal/business/usecases/customeraddressusecase"
	"DB_Project/internal/business/usecases/productusecase"
	"DB_Project/internal/business/usecases/showcaseusecase"
	"database/sql"
)

// All dependencies
var Dependencies Deps

type Deps struct {
	CustomerAddressDeps *CustomerAddressDeps
	ProductDeps         *ProductDeps
	ShowcaseDeps        *ShowcaseDeps
}

type CustomerAddressDeps struct {
	PSQLCustomerAddress customeraddressdomain.CustomerAddressDomain
	PSQLCustomer        customeraddressdomain.CustomerDomain
	PSQLAddress         customeraddressdomain.AddressDomain
}

type ProductDeps struct {
	PSQLProduct productdomain.ProductDomain
}

type ShowcaseDeps struct {
	PSQLShowcase showcasedomain.ShowcaseDomain
}

func InitDeps(db *sql.DB) {
	c := customeraddressusecase.NewPSQLCustomer(db)
	a := customeraddressusecase.NewPSQLAddress(db)

	Dependencies.CustomerAddressDeps = &CustomerAddressDeps{
		PSQLCustomerAddress: customeraddressusecase.NewPSQLCustomerAddress(db, c, a),
		PSQLCustomer:        c,
		PSQLAddress:         a,
	}

	Dependencies.ProductDeps = &ProductDeps{
		PSQLProduct: productusecase.NewPSQLProduct(db),
	}

	Dependencies.ShowcaseDeps = &ShowcaseDeps{
		PSQLShowcase: showcaseusecase.NewPSQLShowcase(db),
	}
}
