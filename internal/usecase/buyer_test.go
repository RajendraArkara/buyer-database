package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/RajendraArkara/buyer-database/internal/entity"
)

type mockBuyerRepo struct {
	findByIDFunc func(ctx context.Context, id int64) (*entity.Buyer, error)
}

func (m *mockBuyerRepo) Create(ctx context.Context, data *entity.Buyer) (int64, error) {
	return 0, nil
}
func (m *mockBuyerRepo) FetchAll(ctx context.Context) ([]entity.Buyer, error) {
	return nil, nil
}
func (m *mockBuyerRepo) FindByID(ctx context.Context, id int64) (*entity.Buyer, error) {
	return m.findByIDFunc(ctx, id)
}
func (m *mockBuyerRepo) UpdateBuyer(ctx context.Context, id int64, data *entity.Buyer) error {
	return nil
}
func (m *mockBuyerRepo) DeleteBuyer(ctx context.Context, id int64) error {
	return nil
}

func TestFindByID(t *testing.T) {
	tests := []struct {
		name      string
		mockBuyer *entity.Buyer
		mockErr   error
		wantErr   bool
	}{
		{
			name:      "berhasil - buyer ketemu",
			mockBuyer: &entity.Buyer{BuyerID: 1, NamaPerusahaan: "PT Contoh"},
			wantErr:   false,
		},
		{
			name:    "gagal - buyer gak ketemu",
			mockErr: errors.New("not found"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockBuyerRepo{
				findByIDFunc: func(ctx context.Context, id int64) (*entity.Buyer, error) {
					return tt.mockBuyer, tt.mockErr
				},
			}
			uc := NewBuyerRepository(repo)

			_, err := uc.FindByID(context.Background(), 1)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
