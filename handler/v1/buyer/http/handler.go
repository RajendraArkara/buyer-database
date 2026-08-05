package http

import (
	"net/http"
	"strconv"

	"github.com/RajendraArkara/buyer-database/internal/usecase"
	"github.com/gin-gonic/gin"
)

type BuyerHandler struct {
	Uc usecase.IBuyerUseCase
}

func NewHandler(uc usecase.IBuyerUseCase) *BuyerHandler {
	return &BuyerHandler{
		Uc: uc,
	}
}

func (h *BuyerHandler) GetAllBuyer(ctx *gin.Context) {
	data, err := h.Uc.FetchAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not fetch the data",
		})
		return
	}

	resp := make([]BuyerObject, 0, len(data))
	for _, item := range data {
		resp = append(resp, BuyerObject{}.ParseFromEntity(item))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func (h *BuyerHandler) GetByID(ctx *gin.Context) {
	buyerid, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Buyer not found!",
		})
		return
	}

	buyer, err := h.Uc.FindByID(ctx, buyerid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the data",
		})
		return
	}

	resp := BuyerObject{}.ParseFromEntity(*buyer)

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func (h *BuyerHandler) CreateBuyer(ctx *gin.Context) {
	var req CreateBuyerRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the data!",
			"error":   err.Error(),
		})
		return
	}

	buyer := req.ToEntity()
	buyer.UserID = 1

	id, err := h.Uc.Create(ctx.Request.Context(), buyer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the data, try again later!",
		})
		return
	}

	buyer.BuyerID = id
	parse := BuyerObject{}.ParseFromEntity(*buyer)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "buyer created!",
		"buyer":   parse,
	})
}

func (h *BuyerHandler) UpdateBuyer(ctx *gin.Context) {
	var req CreateBuyerRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not fetch the data!",
		})
		return
	}

	buyerID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "BuyerID not found!",
		})
		return
	}

	buyer := req.ToEntity()
	buyer.UserID = 1

	err = h.Uc.UpdateBuyer(ctx.Request.Context(), buyerID, buyer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the data!",
		})
		return
	}

	parse := BuyerObject{}.ParseFromEntity(*buyer)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"Detail":  parse,
	})
}
