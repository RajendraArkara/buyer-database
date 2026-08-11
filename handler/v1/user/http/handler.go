package http

import (
	"net/http"

	"github.com/RajendraArkara/buyer-database/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Uc usecase.IUserUseCase
}

func NewHandler(uc usecase.IUserUseCase) *UserHandler {
	return &UserHandler{
		Uc: uc,
	}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var req CreateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data!",
		})
		return
	}

	user := req.ToEntity()

	UserID, err := h.Uc.SignUp(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the data, try again later",
			"error":   err.Error(),
		})
		return
	}

	user.UserID = UserID
	parse := UserObject{}.ParseFromEntity(*user)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created!",
		"User":    parse,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data!",
			"error":   err.Error(),
		})
		return
	}

	token, err := h.Uc.Login(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user!",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
