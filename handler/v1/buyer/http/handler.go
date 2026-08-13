package http

import (
	"database/sql"
	"errors"
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
			"message": "can not fetch the data",
			"err":     err.Error(),
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid buyer id",
			"error":   err.Error(),
		})
		return
	}

	buyer, err := h.Uc.FindByID(ctx, buyerid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "buyer not found",
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch the data",
			"error":   err.Error(),
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
			"message": "could not parse the data",
			"error":   err.Error(),
		})
		return
	}

	buyer := req.ToEntity()
	buyer.UserID = 1

	id, err := h.Uc.Create(ctx.Request.Context(), buyer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch the data, try again later",
			"error":   err.Error(),
		})
		return
	}

	buyer.BuyerID = id
	parse := BuyerObject{}.ParseFromEntity(*buyer)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "buyer created",
		"buyer":   parse,
	})
}

func (h *BuyerHandler) UpdateBuyer(ctx *gin.Context) {
	var req CreateBuyerRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not fetch the data",
			"error":   err.Error(),
		})
		return
	}

	buyerID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "buyerID not found",
			"error":   err.Error(),
		})
		return
	}

	buyer := req.ToEntity()
	buyer.UserID = 1

	err = h.Uc.UpdateBuyer(ctx.Request.Context(), buyerID, buyer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch the data",
			"error":   err.Error(),
		})
		return
	}

	parse := BuyerObject{}.ParseFromEntity(*buyer)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"Detail":  parse,
	})
}

func (h *BuyerHandler) DeleteBuyer(ctx *gin.Context) {
	BuyerID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "ID not found",
			"error":   err.Error(),
		})
		return
	}

	err = h.Uc.DeleteBuyer(ctx.Request.Context(), BuyerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch the data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
