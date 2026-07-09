package entity

import "time"

type Buyer struct {
	BuyerID           int64
	NamaPerusahaan    string `binding:"required"`
	WebsitePerusahaan string `binding:"required"`
	NomorTelepon      string `binding:"required"`
	Negara            string `binding:"required"`
	EmailPerusahaan   string `binding:"required"`
	DateTime          time.Time
	UserID            int64
}

var buyer = []Buyer{}

func (b Buyer) Save() {
	buyer = append(buyer, b)
}

func GetAllBuyer() []Buyer {
	return buyer
}
