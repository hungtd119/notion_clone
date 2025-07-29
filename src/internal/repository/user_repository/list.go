package user_repository

import (
	"context"
	"notion/src/internal/domain/models"
)

// List lấy danh sách user với pagination
func (r *repository) List(ctx context.Context, limit, offset int) ([]models.User, error) {
	var users []models.User
	err := r.GetDB(ctx).Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
