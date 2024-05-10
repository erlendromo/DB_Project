package customeraddressdomain

import "context"

type DBCustomer struct {
	ID          int    `db:"id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	Role        int    `db:"role"`
}

type CreateCustomer struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type CustomerDomain interface {
	CreateCustomer(ctx context.Context, customer *CreateCustomer) (int, error)
	GetCustomerByID(ctx context.Context, id int) (*DBCustomer, error)
	GetCustomerByUsername(ctx context.Context, username string) (*DBCustomer, error)
	UpdateCustomer(ctx context.Context, id int, customer *CreateCustomer) error
	SoftDeleteCustomer(ctx context.Context, id int) error
	GetAdminByUsername(ctx context.Context, username string) (bool, error)
}
