package user_repository

import (
	"context"
	"notion/src/internal/domain/models"
)

// GetByID lấy user theo ID
func (r *repository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.GetDB(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
