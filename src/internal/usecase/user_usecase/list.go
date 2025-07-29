package user_usecase

import (
	"context"
	"notion/src/internal/domain/models"
)

// List lấy danh sách user
func (u *usecase) List(ctx context.Context, limit, offset int) ([]models.User, error) {
	return u.userRepo.List(ctx, limit, offset)
}
