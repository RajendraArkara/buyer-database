package http

import (
	"time"

	"github.com/RajendraArkara/buyer-database/internal/entity"
)

type UserObject struct {
	UserID    int64     `json:"user_id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserObject) ParseFromEntity(e entity.User) UserObject {
	return UserObject{
		UserID:    e.UserID,
		UserName:  e.UserName,
		Email:     e.Email,
		Role:      e.Role,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdateAt,
	}
}
