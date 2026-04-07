package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/repository"
)

// ErrInvalidCredentials is returned when email/password do not match.
var ErrInvalidCredentials = errors.New("invalid credentials")

// ValidateCredentialsUseCase verifies that the given email and password are correct.
type ValidateCredentialsUseCase struct {
	repo repository.UserRepository
}

func NewValidateCredentialsUseCase(repo repository.UserRepository) *ValidateCredentialsUseCase {
	return &ValidateCredentialsUseCase{repo: repo}
}

func (uc *ValidateCredentialsUseCase) Execute(email, password string) (*model.User, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
