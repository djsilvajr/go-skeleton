package usecase

import (
	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
)

// ListUsersUseCase retrieves all users from the repository.
type ListUsersUseCase struct {
	repo repository.UserRepository
}

func NewListUsersUseCase(repo repository.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

func (uc *ListUsersUseCase) Execute() ([]model.User, error) {
	return uc.repo.FindAll()
}
