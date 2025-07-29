package user_usecase

import (
	"context"
	"notion/src/internal/domain/models"
)

// GetByID lấy user theo ID
func (u *usecase) GetByID(ctx context.Context, id uint) (*models.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
