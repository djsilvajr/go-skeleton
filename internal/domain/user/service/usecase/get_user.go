package usecase

import (
	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
)

// GetUserUseCase retrieves a single user by ID.
type GetUserUseCase struct {
	repo repository.UserRepository
}

func NewGetUserUseCase(repo repository.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

func (uc *GetUserUseCase) Execute(id uint) (*model.User, error) {
	return uc.repo.FindByID(id)
}
