package user_usecase

import (
	"context"
	"errors"
	"notion/src/internal/domain/models"

	"golang.org/x/crypto/bcrypt"
)

// Register đăng ký user mới
func (u *usecase) Register(ctx context.Context, user *models.User) error {
	// Kiểm tra email đã tồn tại chưa
	existingUser, err := u.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Tạo user
	return u.userRepo.Create(ctx, user)
}
