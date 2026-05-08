package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	userError "github.com/djsilvajr/go-skeleton/internal/domain/user/errors"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
)

// ValidateCredentialsUseCase verifies that the given email and password are correct.
type ValidateCredentialsUseCase struct {
	repo repository.UserRepository
}

func NewValidateCredentialsUseCase(repo repository.UserRepository) *ValidateCredentialsUseCase {
	return &ValidateCredentialsUseCase{repo: repo}
}

func (uc *ValidateCredentialsUseCase) Execute(email, password string) (*model.User, error) {
	userFound, err := uc.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userError.ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(password)); err != nil {
		return nil, userError.ErrInvalidCredentials
	}

	return userFound, nil
}
