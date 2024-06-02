package domain

type Category struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	MinPrice int    `json:"min_price"`
	MaxPrice int    `json:"max_price"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser(id string) (User, error)
}
