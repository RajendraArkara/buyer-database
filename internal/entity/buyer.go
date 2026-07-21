package entity

import (
	"time"
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
