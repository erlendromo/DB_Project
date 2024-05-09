package productdomain

type Product struct {
	ID               int     `json:"id"`
	CategoryName     string  `json:"category_name"`
	ManufacturerName string  `json:"manufacturer_name"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	Stock            int     `json:"stock"`
}

type ProductDomain interface {
	GetAllProducts() ([]*Product, error)
}
