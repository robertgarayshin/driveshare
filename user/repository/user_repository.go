package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"user-service/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	r.db.Create(&user)
	return user, nil
}

func (r *userRepository) GetUser(id string) (domain.User, error) {
	var user domain.User
	r.db.First(&user, "id = ?", id)
	return user, nil
}
