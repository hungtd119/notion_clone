package user_repository

import (
	"context"
	"notion/src/internal/domain/models"
)

// GetByEmail lấy user theo email
func (r *repository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.GetDB(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
