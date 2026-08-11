package entity

import "time"

type User struct {
	UserID    int64
	UserName  string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	Role      string `binding:"required"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
