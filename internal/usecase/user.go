package usecase

import (
	"context"
	"errors"

	"github.com/RajendraArkara/buyer-database/infrastructure/utils"
	"github.com/RajendraArkara/buyer-database/internal/entity"
	"github.com/RajendraArkara/buyer-database/internal/repository"
)

type IUserUseCase interface {
	SignUp(ctx context.Context, data *entity.User) (int64, error)
	Login(ctx context.Context, email, pass string) (string, error)
}

type UserUseCase struct {
	Repo repository.UserRepository
}

func NewUserRepository(repo repository.UserRepository) IUserUseCase {
	return &UserUseCase{
		Repo: repo,
	}
}

func (uc *UserUseCase) SignUp(ctx context.Context, data *entity.User) (int64, error) {
	passHash, err := utils.HashPassword(data.Password)
	if err != nil {
		return 0, err
	}

	data.Password = passHash

	id, err := uc.Repo.SignUp(ctx, data)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email, pass string) (string, error) {
	user, err := uc.Repo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Credentials invalid!")
	}

	if !utils.CheckPasswordHash(pass, user.Password) {
		return "", errors.New("Credentials invalid!")
	}

	return utils.GenerateToken(user.Email, user.UserID)
}
