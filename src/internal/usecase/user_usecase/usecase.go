package user_usecase

import (
	"notion/src/internal/domain/repositories"
	"notion/src/internal/domain/usecases"
)

type usecase struct {
	userRepo repositories.UserRepository
}

func New() usecases.UserUseCase {
	return &usecase{}
}

func (u *usecase) SetUserRepo(userRepo repositories.UserRepository) {
	u.userRepo = userRepo
}
