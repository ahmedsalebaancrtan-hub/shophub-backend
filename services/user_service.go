
package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/ahmed/shophub-backend/dto"
	"github.com/ahmed/shophub-backend/models"
	"github.com/ahmed/shophub-backend/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

type UserService struct {
	repo *repository.UserRepo
}

func RegisterUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) CreateUser(
	ctx context.Context,
	data dto.RegisterRequest,
) (dto.RegisterResponse, error) {

	email := strings.ToLower(strings.TrimSpace(data.Email))
	fullName := strings.TrimSpace(data.FullName)

	// Check whether the email already exists.
	_, err := svc.repo.GetUserbyEmail(ctx, email)

	if err == nil {
		return dto.RegisterResponse{}, ErrEmailAlreadyExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.RegisterResponse{}, err
	}

	// Hash the password.
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// Prepare user model.
	user := models.User{
		FullName:     fullName,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         "CUSTOMER",
	}

	// Save user in database.
	if err := svc.repo.CreateUser(ctx, &user); err != nil {
		return dto.RegisterResponse{}, err
	}

	// Return safe response.
	response := dto.RegisterResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	return response, nil
}
