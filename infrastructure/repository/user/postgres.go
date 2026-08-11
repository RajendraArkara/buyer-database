package user

import (
	"context"
	"database/sql"

	"github.com/RajendraArkara/buyer-database/internal/entity"
	"github.com/RajendraArkara/buyer-database/internal/repository"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) SignUp(ctx context.Context, data *entity.User) (int64, error) {
	query := `
		INSERT INTO users (user_name, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id, created_at
	`
	var id int64

	err := r.db.QueryRow(
		query,
		data.UserName,
		data.Email,
		data.Password,
		data.Role,
	).Scan(&id, &data.CreatedAt)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) { // masih bingung soal ini
	query := `
		SELECT user_id, email, password
		FROM users
		WHERE email = $1
	`
	var user entity.User

	err := r.db.QueryRow(query, email).Scan(&user.UserID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
