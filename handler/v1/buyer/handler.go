package buyer

import (
	"net/http"

	"github.com/RajendraArkara/buyer-database/internal/entity"
	"github.com/gin-gonic/gin"
)

func GetAllBuyer(ctx *gin.Context) {
	buyer, err := entity.GetAllBuyer()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not fetch data",
		})
		return
	}
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

	buyer.UserID = 1

	id, err := buyer.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the data, try again later!",
		})
		return
	}

	buyer.UserID = id

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "buyer created!",
		"buyer":   buyer,
	})
}
