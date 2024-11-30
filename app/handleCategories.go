package app

import (
	"jellybar/db"
	"jellybar/obj"

	"github.com/HazelnutParadise/Go-Utils/conv"
	"github.com/gin-gonic/gin"
)

func HandleGetCategories(ctx *gin.Context) {
	preloadArticles := ctx.Query("preloadArticles") == "true"
	categories, err := db.GetCategories(preloadArticles)
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

func HandleDeleteCategory(ctx *gin.Context) {
	id := uint(conv.ParseInt(ctx.Param("id")))
	if id == 0 {
		ctx.JSON(400, gin.H{"message": "刪除類別失敗\nID錯誤"})
		return
	}
	category := obj.Category{ID: id}
	if err := db.DeleteCategory(&category); err != nil {
		ctx.JSON(500, gin.H{"message": "刪除類別失敗\n" + err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "刪除類別成功"})
}
