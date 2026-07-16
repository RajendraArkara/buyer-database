package handler

import (
	"github.com/RajendraArkara/buyer-database/handler/v1/buyer"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.GET("/buyer", buyer.GetAllBuyer)
	server.POST("/create-buyer", buyer.CreateBuyer)
	server.GET("/buyer/:id", buyer.GetByID)
}
