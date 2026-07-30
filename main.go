package main

import (
	"github.com/RajendraArkara/buyer-database/handler"
	buyerhttp "github.com/RajendraArkara/buyer-database/handler/v1/buyer/http"
	"github.com/RajendraArkara/buyer-database/infrastructure/db"
	buyerrepo "github.com/RajendraArkara/buyer-database/infrastructure/repository/buyer"
	"github.com/RajendraArkara/buyer-database/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	buyerRepo := buyerrepo.NewBuyerRepository(db.DB)
	buyerUseCase := usecase.NewBuyerRepository(buyerRepo)
	buyerHandler := buyerhttp.NewHandler(buyerUseCase)

	server := gin.Default()
	handler.Routes(server, buyerHandler)

	server.Run(":8080")
}
