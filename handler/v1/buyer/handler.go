package buyer

import (
	"net/http"

	"github.com/RajendraArkara/buyer-database/internal/entity"
	"github.com/gin-gonic/gin"
)

func GetAllBuyer(ctx *gin.Context) {
	buyer := entity.GetAllBuyer()
	ctx.JSON(http.StatusOK, buyer)
}

func Create(ctx *gin.Context) {}
