package buyer

import (
	"context"
	"database/sql"

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
		RETURNING buyer_id, date_time
	`

	var id int64

	err := r.db.QueryRow(query,
		data.NamaPerusahaan,
		data.WebsitePerusahaan,
		data.NomorTelepon,
		data.EmailPerusahaan,
		data.Negara,
		data.KomoditasPerusahaan,
		data.UserID,
	).Scan(&id, &data.DateTime)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresRepository) FetchAll(ctx context.Context) ([]entity.Buyer, error) {
	query := `
		SELECT * FROM buyers
	`
	rows, err := r.db.Query(query)

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

	row := r.db.QueryRow(query, id)

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

func (r *PostgresRepository) UpdateBuyer(ctx context.Context, id int64, data *entity.Buyer) error {
	query := `
		UPDATE buyers
		SET company_name = $1, company_website = $2, company_telephone = $3, company_email = $4, country = $5, commodity = $6
		WHERE buyer_id = $7
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		&data.NamaPerusahaan,
		&data.WebsitePerusahaan,
		&data.NomorTelepon,
		&data.EmailPerusahaan,
		&data.Negara,
		&data.KomoditasPerusahaan,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
