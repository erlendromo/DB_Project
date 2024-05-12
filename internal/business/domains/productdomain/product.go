package productdomain

type Product struct {
	ID               int     `json:"id"`
	CategoryName     string  `json:"category_name"`
	ManufacturerName string  `json:"manufacturer_name"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	Stock            int     `json:"stock"`
}

type PointerProduct struct {
	ID               *int     `json:"id"`
	CategoryName     *string  `json:"category_name"`
	ManufacturerName *string  `json:"manufacturer_name"`
	Description      *string  `json:"description"`
	Price            *float64 `json:"price"`
	Stock            *int     `json:"stock"`
}

type ProductDetail struct {
	ID               int      `json:"id"`
	CategoryName     string   `json:"category_name"`
	ManufacturerName string   `json:"manufacturer_name"`
	Description      string   `json:"description"`
	Price            float64  `json:"price"`
	Stock            int      `json:"stock"`
	AverageRating    float64  `json:"average_rating,omitempty"`
	Reviews          []Review `json:"reviews,omitempty"`
}

type Review struct {
	CustomerID int     `json:"customer_id"`
	Stars      float64 `json:"stars"`
	Comment    string  `json:"comment"`
}

type ProductDomain interface {
	GetAllProducts() ([]*Product, error)
	GetProduct(id string) (*ProductDetail, error)
	PostProduct(product *Product) (int, error)
	SearchProductFullText(description string) ([]*Product, error)
	DeleteProduct(id string) error
	PatchProduct(id string, product *PointerProduct) (*Product, error)
}
