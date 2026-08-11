package main

import (
	"github.com/RajendraArkara/buyer-database/handler"
	buyerhttp "github.com/RajendraArkara/buyer-database/handler/v1/buyer/http"
	userhttp "github.com/RajendraArkara/buyer-database/handler/v1/user/http"
	"github.com/RajendraArkara/buyer-database/infrastructure/db"
	buyerrepo "github.com/RajendraArkara/buyer-database/infrastructure/repository/buyer"
	userrepo "github.com/RajendraArkara/buyer-database/infrastructure/repository/user"
	"github.com/RajendraArkara/buyer-database/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	//buyer
	buyerRepo := buyerrepo.NewBuyerRepository(db.DB)
	buyerUseCase := usecase.NewBuyerRepository(buyerRepo)
	buyerHandler := buyerhttp.NewHandler(buyerUseCase)

	//user
	userRepo := userrepo.NewUserRepository(db.DB)
	userUseCse := usecase.NewUserRepository(userRepo)
	userHandler := userhttp.NewHandler(userUseCse)

	server := gin.Default()
	handler.RoutesBuyer(server, buyerHandler)
	handler.RoutesUser(server, userHandler)

	server.Run(":8080")
}
