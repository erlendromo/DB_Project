package userdomain

import "context"

type UserDomain interface {
	Insert(ctx context.Context, user *User) error
	SelectByID(ctx context.Context, id int) (User, error)
	SelectAll(ctx context.Context) ([]User, error)
}

type User struct {
	ID          int    `psql:"id,omitempty"`
	Username    string `psql:"username"`
	Password    string `psql:"password"`
	FirstName   string `psql:"first_name"`
	LastName    string `psql:"last_name"`
	Email       string `psql:"email"`
	PhoneNumber string `psql:"phone_number"`
	Role        int16  `psql:"role,omitempty"`
}
