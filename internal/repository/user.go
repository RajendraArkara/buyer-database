package repository

import (
	"context"

	"github.com/RajendraArkara/buyer-database/internal/entity"
)

type UserRepository interface {
	SignUp(ctx context.Context, data *entity.User) (int64, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
