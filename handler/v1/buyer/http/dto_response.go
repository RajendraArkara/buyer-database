package http

import (
	"time"

	"github.com/RajendraArkara/buyer-database/internal/entity"
)

type BuyerObject struct {
	BuyerID             int64     `json:"buyer_id"`
	NamaPerusahaan      string    `json:"nama_perusahaan"`
	WebsitePerusahaan   string    `json:"website_perusahaan"`
	NomorTelepon        string    `json:"nomor_telepon"`
	Negara              string    `json:"negara"`
	EmailPerusahaan     string    `json:"email_perusahaan"`
	KomoditasPerusahaan string    `json:"komoditas_perusahaan"`
	DateTime            time.Time `json:"waktu"`
	UserID              int64     `json:"user_id"`
}

func (BuyerObject) ParseFromEntity(e entity.Buyer) BuyerObject {
	return BuyerObject{
		BuyerID:             e.BuyerID,
		NamaPerusahaan:      e.NamaPerusahaan,
		WebsitePerusahaan:   e.WebsitePerusahaan,
		NomorTelepon:        e.NomorTelepon,
		Negara:              e.Negara,
		EmailPerusahaan:     e.EmailPerusahaan,
		KomoditasPerusahaan: e.KomoditasPerusahaan,
		DateTime:            e.DateTime,
		UserID:              e.UserID,
	}
}
