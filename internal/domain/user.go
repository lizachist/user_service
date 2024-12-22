package domain

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"-"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	Update(user *User) error
	GetByUsername(username string) (*User, error)
}

type UserService interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	Update(user *User) error
	Authenticate(username, password string) (*User, error)
}
