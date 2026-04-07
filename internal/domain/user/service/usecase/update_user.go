package usecase

import (
	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
)

// UpdateUserInput carries the fields that can be updated.
type UpdateUserInput struct {
	ID    uint
	Name  string
	Email string
}

// UpdateUserUseCase contains the business rules for updating a user.
type UpdateUserUseCase struct {
	repo repository.UserRepository
}

func NewUpdateUserUseCase(repo repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInput) (*model.User, error) {
	user, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := uc.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
