package http

import "github.com/RajendraArkara/buyer-database/internal/entity"

type CreateBuyerRequest struct {
	NamaPerusahaan      string `json:"nama_perusahaan" binding:"required"`
	WebsitePerusahaan   string `json:"website_perusahaan" binding:"required"`
	NomorTelepon        string `json:"nomor_telepon" binding:"required"`
	Negara              string `json:"negara" binding:"required"`
	EmailPerusahaan     string `json:"email_perusahaan" binding:"required"`
	KomoditasPerusahaan string `json:"komoditas_perusahaan" binding:"required"`
}

func (dto CreateBuyerRequest) ToEntity() *entity.Buyer {
	return &entity.Buyer{
		NamaPerusahaan:      dto.NamaPerusahaan,
		WebsitePerusahaan:   dto.WebsitePerusahaan,
		NomorTelepon:        dto.NomorTelepon,
		Negara:              dto.Negara,
		EmailPerusahaan:     dto.EmailPerusahaan,
		KomoditasPerusahaan: dto.KomoditasPerusahaan,
	}
}
