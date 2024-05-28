package domain

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser(id string) (User, error)
}
