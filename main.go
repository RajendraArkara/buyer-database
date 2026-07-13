package main

import (
	"github.com/RajendraArkara/buyer-database/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	handler.Routes(server)

	server.Run(":8080")
}
