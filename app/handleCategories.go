package app

import (
	"jellybar/db"
	"jellybar/obj"

	"github.com/gin-gonic/gin"
)

func HandleGetCategories(ctx *gin.Context) {
	categories, err := db.GetCategories(false)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "取得類別列表失敗\n" + err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"categories": categories})
}

func HandlePostCategory(ctx *gin.Context) {
	var category obj.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{"message": "新增類別失敗\n" + err.Error()})
		return
	}
	if err := db.AddCategory(&category); err != nil {
		ctx.JSON(500, gin.H{"message": "新增類別失敗\n" + err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "新增類別成功"})
}
