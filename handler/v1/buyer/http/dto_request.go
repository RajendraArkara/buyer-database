package http

type CreateBuyerRequest struct {
	NamaPerusahaan string `json:"nama_perusahaan" binding:"required"`
}
