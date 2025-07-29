package user_repository

import (
	"context"
	"notion/src/internal/domain/models"
)

// Update cập nhật user
func (r *repository) Update(ctx context.Context, user *models.User) error {
	return r.GetDB(ctx).Save(user).Error
}
