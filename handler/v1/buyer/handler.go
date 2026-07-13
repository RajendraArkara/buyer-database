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

func CreateBuyer(ctx *gin.Context) {
	var buyer entity.Buyer
	err := ctx.ShouldBindJSON(&buyer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the data!",
		})
		return
	}

	buyer.BuyerID = 1
	buyer.UserID = 1

	buyer.Save()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "buyer created!",
		"buyer":   buyer,
	})
}
