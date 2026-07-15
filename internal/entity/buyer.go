package entity

import (
	"time"

	"github.com/RajendraArkara/buyer-database/infrastructure/db"
)

type Buyer struct {
	BuyerID             int64
	NamaPerusahaan      string `binding:"required"`
	WebsitePerusahaan   string `binding:"required"`
	NomorTelepon        string `binding:"required"`
	Negara              string `binding:"required"`
	EmailPerusahaan     string `binding:"required"`
	KomoditasPerusahaan string `binding:"required"`
	DateTime            time.Time
	UserID              int64
}

// var buyer = []Buyer{}

func (b Buyer) Save() (int64, error) {
	query := `
		INSERT INTO buyers (company_name, company_website, company_telephone, company_email, country, commodity, user_id)
		VALUES ($1, $2 , $3, $4, $5, $6, $7)
		RETURNING buyer_id
	`
	var id int64

	err := db.DB.QueryRow(query,
		b.NamaPerusahaan,
		b.WebsitePerusahaan,
		b.NomorTelepon,
		b.EmailPerusahaan,
		b.Negara,
		b.KomoditasPerusahaan,
		b.UserID,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetAllBuyer() ([]Buyer, error) {
	query := `
		SELECT * FROM buyers
	`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var buyers []Buyer

	for rows.Next() {
		var buyer Buyer

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
