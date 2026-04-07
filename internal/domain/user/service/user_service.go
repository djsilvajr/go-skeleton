package service

import (
	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/service/usecase"
)

// UserService defines user business operations.
type UserService interface {
	List() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	Create(name, email, password string) (*model.User, error)
	Update(id uint, name, email string) (*model.User, error)
	Delete(id uint) error
	ValidateCredentials(email, password string) (*model.User, error)
}

// userService orchestrates use cases — it contains no business logic itself.
type userService struct {
	listUsers           *usecase.ListUsersUseCase
	getUser             *usecase.GetUserUseCase
	createUser          *usecase.CreateUserUseCase
	updateUser          *usecase.UpdateUserUseCase
	deleteUser          *usecase.DeleteUserUseCase
	validateCredentials *usecase.ValidateCredentialsUseCase
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		listUsers:           usecase.NewListUsersUseCase(repo),
		getUser:             usecase.NewGetUserUseCase(repo),
		createUser:          usecase.NewCreateUserUseCase(repo),
		updateUser:          usecase.NewUpdateUserUseCase(repo),
		deleteUser:          usecase.NewDeleteUserUseCase(repo),
		validateCredentials: usecase.NewValidateCredentialsUseCase(repo),
	}
}

func (s *userService) List() ([]model.User, error) {
	return s.listUsers.Execute()
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.getUser.Execute(id)
}

func (s *userService) Create(name, email, password string) (*model.User, error) {
	return s.createUser.Execute(usecase.CreateUserInput{
		Name:     name,
		Email:    email,
		Password: password,
	})
}

func (s *userService) Update(id uint, name, email string) (*model.User, error) {
	return s.updateUser.Execute(usecase.UpdateUserInput{
		ID:    id,
		Name:  name,
		Email: email,
	})
}

func (s *userService) Delete(id uint) error {
	return s.deleteUser.Execute(id)
}

func (s *userService) ValidateCredentials(email, password string) (*model.User, error) {
	return s.validateCredentials.Execute(email, password)
}
