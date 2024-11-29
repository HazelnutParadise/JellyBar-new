package app

import (
	"jellybar/db"

	"github.com/gin-gonic/gin"
)

func HandleGetCategories(ctx *gin.Context) {
	categories, err := db.GetCategories()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "取得類別列表失敗\n" + err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"categories": categories})
}
