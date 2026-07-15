package main

import (
	"github.com/RajendraArkara/buyer-database/handler"
	"github.com/RajendraArkara/buyer-database/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	handler.Routes(server)

	server.Run(":8080")
}
