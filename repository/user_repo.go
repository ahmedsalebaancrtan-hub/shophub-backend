package repository

import (
	"context"

	"github.com/ahmed/shophub-backend/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func RegisterUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo UserRepo) GetUserbyEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := repo.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepo) CreateUser(ctx context.Context,user *models.User) error {
	return repo.DB.WithContext(ctx).Create(user).Error
}
