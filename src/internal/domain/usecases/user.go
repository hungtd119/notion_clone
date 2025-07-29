package usecases

import (
	"context"
	"notion/src/internal/domain/models"
	"notion/src/internal/domain/repositories"
)

// UserUseCase interface định nghĩa business logic
type UserUseCase interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email, password string) (*models.User, error)
	GetByID(ctx context.Context, id uint) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]models.User, error)

	SetUserRepo(userRepo repositories.UserRepository)
}
