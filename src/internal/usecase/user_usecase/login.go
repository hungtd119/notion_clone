package user_usecase

import (
	"context"
	"errors"
	"notion/src/internal/domain/models"

	"golang.org/x/crypto/bcrypt"
)

// Login đăng nhập user
func (u *usecase) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Kiểm tra password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
