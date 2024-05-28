package usecase

import "user/domain"

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) *UserUsecase {
	return &UserUsecase{r}
}

func (u *UserUsecase) CreateUser(name, email, password string) (domain.User, error) {
	user := domain.User{
		ID:       generateID(),
		Name:     name,
		Email:    email,
		Password: password,
	}
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(id string) (domain.User, error) {
	return u.userRepo.GetUser(id)
}
