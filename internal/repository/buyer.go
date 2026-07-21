package repository

import (
	"context"

	"github.com/RajendraArkara/buyer-database/internal/entity"
)

type BuyerRepository interface {
	Create(ctx context.Context, data *entity.Buyer) (int64, error)
	FetchAll(ctx context.Context) ([]entity.Buyer, error)
	FindByID(ctx context.Context, id int64) (*entity.Buyer, error)
}
