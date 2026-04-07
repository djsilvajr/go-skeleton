package usecase

import "github.com/djsilvajr/go-skeleton/internal/domain/user/repository"

// DeleteUserUseCase performs a soft-delete of the user.
// The repository layer delegates to GORM which sets the deleted_at column
// without permanently removing the record, satisfying the soft-delete rule.
type DeleteUserUseCase struct {
	repo repository.UserRepository
}

func NewDeleteUserUseCase(repo repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Execute(id uint) error {
	return uc.repo.Delete(id)
}
