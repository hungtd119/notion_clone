package user_usecase

import (
	"context"
	"notion/src/internal/domain/models"

	"golang.org/x/crypto/bcrypt"
)

// Update cập nhật user
func (u *usecase) Update(ctx context.Context, user *models.User) error {
	// Nếu có password mới, hash nó
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return u.userRepo.Update(ctx, user)
}
