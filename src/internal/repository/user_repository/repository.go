package user_repository

import (
	"notion/src/internal/domain/repositories"
	"notion/src/internal/repository/base_repository"
)

type repository struct {
	base_repository.BaseRepository
}

func New() repositories.UserRepository {
	return &repository{}
}
