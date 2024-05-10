package dependencies

import (
	"DB_Project/internal/business/domains/customeraddressdomain"
	"DB_Project/internal/business/domains/productdomain"
	"DB_Project/internal/business/usecases/customeraddressusecase"
	"DB_Project/internal/business/usecases/productusecase"
	"DB_Project/internal/business/usecases/showcase"
	"database/sql"
)

var Dependencies Deps

type Deps struct {
	CustomerDeps *CustomerDeps
	ProductDeps  *ProductDeps
	ShowcaseDeps *ShowcaseDeps
}

func GetDeps() *Deps {
	return &Dependencies
}

func (d *Deps) GetCustomerDeps() *CustomerDeps {
	return d.CustomerDeps
}

func InitDeps(db *sql.DB) {
	Dependencies.CustomerDeps = newCustomerDeps(
		customeraddressusecase.NewPSQLCustomer(db),
		customeraddressusecase.NewPSQLAddress(db),
		customeraddressusecase.NewPSQLCustomerAddress(db),
	)
	Dependencies.ProductDeps = newProductDeps(
		productusecase.NewPSQLProduct(db),
	)
	Dependencies.ShowcaseDeps = newShowcaseDeps(
		showcase.NewPSQLShowcase(db),
	)
}

type CustomerDeps struct {
	CustomerDomain        customeraddressdomain.CustomerDomain
	AddressDomain         customeraddressdomain.AddressDomain
	CustomerAddressDomain customeraddressdomain.CustomerAddressDomain
}

func (cd *CustomerDeps) GetCustomerDomain() customeraddressdomain.CustomerDomain {
	return cd.CustomerDomain
}

func (cd *CustomerDeps) GetAddressDomain() customeraddressdomain.AddressDomain {
	return cd.AddressDomain
}

func (cd *CustomerDeps) GetCustomerAddressDomain() customeraddressdomain.CustomerAddressDomain {
	return cd.CustomerAddressDomain
}

func newCustomerDeps(cd customeraddressdomain.CustomerDomain, ad customeraddressdomain.AddressDomain, cad customeraddressdomain.CustomerAddressDomain) *CustomerDeps {
	return &CustomerDeps{
		CustomerDomain:        cd,
		AddressDomain:         ad,
		CustomerAddressDomain: cad,
	}
}

type ProductDeps struct {
	ProductDomain productdomain.ProductDomain
}

func newProductDeps(pd productdomain.ProductDomain) *ProductDeps {
	return &ProductDeps{
		ProductDomain: pd,
	}
}

type ShowcaseDeps struct {
	ShowcaseDomain showcase.Domain
}

func newShowcaseDeps(sd showcase.Domain) *ShowcaseDeps {
	return &ShowcaseDeps{
		ShowcaseDomain: sd,
	}
}
