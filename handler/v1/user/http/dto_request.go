package http

import "github.com/RajendraArkara/buyer-database/internal/entity"

type CreateUserRequest struct {
	Username string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (dto CreateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		UserName: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     dto.Role,
	}
}
