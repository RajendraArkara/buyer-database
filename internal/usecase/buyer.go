package usecase

import (
	"context"

	"github.com/RajendraArkara/buyer-database/internal/entity"
	"github.com/RajendraArkara/buyer-database/internal/repository"
)

type IBuyerUseCase interface {
	Create(ctx context.Context, data *entity.Buyer) (int64, error)
	FetchAll(ctx context.Context) ([]entity.Buyer, error)
	FindByID(ctx context.Context, id int64) (*entity.Buyer, error)
}

type BuyerUsecase struct {
	Repo repository.BuyerRepository
}

func NewBuyerRepository(repo repository.BuyerRepository) IBuyerUseCase {
	return &BuyerUsecase{
		Repo: repo,
	}
}

func (uc *BuyerUsecase) Create(ctx context.Context, data *entity.Buyer) (int64, error) {
	id, err := uc.Repo.Create(ctx, data)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *BuyerUsecase) FetchAll(ctx context.Context) ([]entity.Buyer, error) {
	data, err := uc.Repo.FetchAll(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *BuyerUsecase) FindByID(ctx context.Context, id int64) (*entity.Buyer, error) {
	data, err := uc.Repo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return data, nil
}
