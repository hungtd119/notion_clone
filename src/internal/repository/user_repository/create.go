package user_repository

import (
	"context"
	"notion/src/internal/domain/models"
)

// Create tạo user mới
func (r *repository) Create(ctx context.Context, user *models.User) error {
	return r.GetDB(ctx).Create(user).Error
}
