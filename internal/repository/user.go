package repository

import (
	"github.com/eigakan/user-service/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *model.User) error {
	return r.db.Create(u).Error
}

func (r *UserRepository) GetUserById(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByLogin(l string) (*model.User, error) {
	var user model.User
	err := r.db.Where("login = ?", l).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
