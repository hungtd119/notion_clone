package user_repository

import (
	"context"
	"notion/src/internal/domain/models"
)

// Delete xóa user theo ID
func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.GetDB(ctx).Delete(&models.User{}, id).Error
}
