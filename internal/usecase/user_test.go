package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/RajendraArkara/buyer-database/infrastructure/utils"
	"github.com/RajendraArkara/buyer-database/internal/entity"
)

type mockUserRepo struct {
	findByEmailFunc func(ctx context.Context, email string) (*entity.User, error)
}

func (m *mockUserRepo) SignUp(ctx context.Context, data *entity.User) (int64, error) {
	return 0, nil
}

func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return m.findByEmailFunc(ctx, email)
}

func (m *mockUserRepo) UpdatePassword(ctx context.Context, id int64, data *entity.User) error {
	return nil
}

func (m *mockUserRepo) GetAll(ctx context.Context) ([]entity.User, error) {
	return nil, nil
}

func TestLogin(t *testing.T) {
	hashedPassword, _ := utils.HashPassword("password-benar")

	tests := []struct {
		name     string
		password string
		mockUser *entity.User
		mockErr  error
		wantErr  bool
	}{
		{
			name:     "berhasil - password benar",
			password: "password-benar",
			mockUser: &entity.User{UserID: 1, Email: "test@example.com", Password: hashedPassword},
			wantErr:  false,
		},
		{
			name:     "gagal - password salah",
			password: "password-salah",
			mockUser: &entity.User{UserID: 1, Email: "test@example.com", Password: hashedPassword},
			wantErr:  true,
		},
		{
			name:     "gagal - user gak ketemu",
			password: "apapun",
			mockUser: nil,
			mockErr:  errors.New("not found"),
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockUserRepo{
				findByEmailFunc: func(ctx context.Context, email string) (*entity.User, error) {
					return tt.mockUser, tt.mockErr
				},
			}
			uc := NewUserRepository(repo)

			_, err := uc.Login(context.Background(), "test@example.com", tt.password)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
