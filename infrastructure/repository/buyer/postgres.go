package buyer

import (
	"context"
	"database/sql"

	"github.com/RajendraArkara/buyer-database/infrastructure/db"
	"github.com/RajendraArkara/buyer-database/internal/entity"
	"github.com/RajendraArkara/buyer-database/internal/repository"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) repository.BuyerRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, data *entity.Buyer) (int64, error) {
	query := `
		INSERT INTO buyers (company_name, company_website, company_telephone, company_email, country, commodity, user_id)
		VALUES ($1, $2 , $3, $4, $5, $6, $7)
		RETURNING buyer_id
	`

	var id int64

	err := db.DB.QueryRow(query,
		data.NamaPerusahaan,
		data.WebsitePerusahaan,
		data.NomorTelepon,
		data.EmailPerusahaan,
		data.Negara,
		data.KomoditasPerusahaan,
		data.UserID,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresRepository) FetchAll(ctx context.Context) ([]entity.Buyer, error) {
	query := `
		SELECT * FROM buyers
	`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var buyers []entity.Buyer

	for rows.Next() {
		var buyer entity.Buyer

		err := rows.Scan(
			&buyer.BuyerID,
			&buyer.NamaPerusahaan,
			&buyer.WebsitePerusahaan,
			&buyer.NomorTelepon,
			&buyer.EmailPerusahaan,
			&buyer.Negara,
			&buyer.DateTime,
			&buyer.UserID,
			&buyer.KomoditasPerusahaan,
		)

		if err != nil {
			return nil, err
		}

		buyers = append(buyers, buyer)
	}

	return buyers, nil
}

func (r *PostgresRepository) FindByID(ctx context.Context, id int64) (*entity.Buyer, error) {
	query := `
		SELECT * FROM buyers
		WHERE buyer_id = $1
	`

	row := db.DB.QueryRow(query, id)

	var buyer entity.Buyer

	err := row.Scan(
		&buyer.BuyerID,
		&buyer.NamaPerusahaan,
		&buyer.WebsitePerusahaan,
		&buyer.NomorTelepon,
		&buyer.EmailPerusahaan,
		&buyer.Negara,
		&buyer.DateTime,
		&buyer.UserID,
		&buyer.KomoditasPerusahaan,
	)

	if err != nil {
		return nil, err
	}

	return &buyer, nil
}
