package handler

import (
	buyerhttp "github.com/RajendraArkara/buyer-database/handler/v1/buyer/http"
	userhttp "github.com/RajendraArkara/buyer-database/handler/v1/user/http"
	"github.com/RajendraArkara/buyer-database/middlewares"
	"github.com/gin-gonic/gin"
)

func RoutesBuyer(server *gin.Engine, h *buyerhttp.BuyerHandler) {
	//buyer
	server.GET("/buyer", h.GetAllBuyer)
	server.POST("/buyer/create-buyer", middlewares.Authencticate, h.CreateBuyer)
	server.GET("/buyer/:id", h.GetByID)
	server.PATCH("/buyer/update/:id", h.UpdateBuyer)
	server.DELETE("/buyer/delete/:id", h.DeleteBuyer)

}

func RoutesUser(server *gin.Engine, h *userhttp.UserHandler) {
	//user
	server.POST("/user/create-user", h.SignUp)
	server.POST("/user/login", h.Login)
	server.PATCH("/user/forgot-password", h.ForgotPassword)
	server.GET("/user/get-all", h.GetAll)
}
