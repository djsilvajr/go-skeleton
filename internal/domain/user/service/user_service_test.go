package service_test

import (
	"errors"
	"testing"

	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
	"github.com/djsilvajr/go-skeleton/internal/domain/user/service"
)

// mockUserRepository is a hand-rolled test double.
// In larger projects replace with mockery or gomock generated mocks.
type mockUserRepository struct {
	users  []model.User
	findFn func(email string) (*model.User, error)
}

func (m *mockUserRepository) FindAll() ([]model.User, error)        { return m.users, nil }
func (m *mockUserRepository) FindByID(id uint) (*model.User, error) { return &m.users[0], nil }
func (m *mockUserRepository) Create(u *model.User) error            { m.users = append(m.users, *u); return nil }
func (m *mockUserRepository) Update(u *model.User) error            { return nil }
func (m *mockUserRepository) Delete(id uint) error                  { return nil }
func (m *mockUserRepository) FindByEmail(email string) (*model.User, error) {
	if m.findFn != nil {
		return m.findFn(email)
	}
	return nil, errors.New("not found")
}

func TestUserService_Create(t *testing.T) {
	repo := &mockUserRepository{}
	svc := service.NewUserService(repo)

	user, err := svc.Create("Alice", "alice@example.com", "secret123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Email != "alice@example.com" {
		t.Errorf("expected email alice@example.com, got %s", user.Email)
	}
	if user.Password == "secret123" {
		t.Error("password must be hashed")
	}
}

func TestUserService_ValidateCredentials_InvalidEmail(t *testing.T) {
	repo := &mockUserRepository{
		findFn: func(email string) (*model.User, error) {
			return nil, errors.New("not found")
		},
	}
	svc := service.NewUserService(repo)

	_, err := svc.ValidateCredentials("unknown@example.com", "pass")
	if err == nil {
		t.Fatal("expected error for unknown user")
	}
}

func TestUserService_List(t *testing.T) {
	repo := &mockUserRepository{
		users: []model.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
	svc := service.NewUserService(repo)

	users, err := svc.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}
