package handler

import (
	buyerhttp "github.com/RajendraArkara/buyer-database/handler/v1/buyer/http"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine, h *buyerhttp.BuyerHandler) {
	server.GET("/buyer", h.GetAllBuyer)
	server.POST("/create-buyer", h.CreateBuyer)
	server.GET("/buyer/:id", h.GetByID)
}
