package usecase

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
)

// CreateUserInput carries the data needed to create a new user.
type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	Role     model.Role
}

// CreateUserUseCase contains the business rules for user creation.
type CreateUserUseCase struct {
	repo repository.UserRepository
}

func NewCreateUserUseCase(repo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

// Execute hashes the password, applies a default role and persists the user.
func (uc *CreateUserUseCase) Execute(input CreateUserInput) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	role := input.Role
	if role == "" {
		role = model.RoleUser
	}

	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Role:     role,
	}

	if err := uc.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
