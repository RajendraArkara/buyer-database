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
			"message": "could not parse the data",
		})
		return
	}

	user := req.ToEntity()

	UserID, err := h.Uc.SignUp(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch the data, try again later",
			"error":   err.Error(),
		})
		return
	}

	user.UserID = UserID
	parse := UserObject{}.ParseFromEntity(*user)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"User":    parse,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the data",
			"error":   err.Error(),
		})
		return
	}

	token, err := h.Uc.Login(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "could not authenticate user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *UserHandler) ForgotPassword(ctx *gin.Context) {
	var req LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the data",
			"error":   err.Error(),
		})
		return
	}

	err = h.Uc.ForgotPassword(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch the data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	data, err := h.Uc.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messasge": "can not fetch the data",
			"error":    err.Error(),
		})
		return
	}

	resp := make([]UserObject, 0, len(data))

	for _, item := range data {
		resp = append(resp, UserObject{}.ParseFromEntity(item))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": resp,
	})
}
