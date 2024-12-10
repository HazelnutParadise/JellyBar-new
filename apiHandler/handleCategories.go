package apiHandler

import (
	"jellybar/db"
	"jellybar/obj"
	"jellybar/utils"

	"github.com/HazelnutParadise/Go-Utils/conv"
	"github.com/gin-gonic/gin"
)

func HandleGetCategories(ctx *gin.Context, isAdminMode bool) {
	preloadArticles := ctx.Query("preloadArticles") == "true"
	var onlyPublished bool
	if isAdminMode {
		onlyPublished = false
	} else {
		onlyPublished = true
	}
	categories, err := db.GetCategories(db.GetCategoriesOption{
		PreloadArticles: preloadArticles,
		OnlyPublished:   onlyPublished,
	})
	if err != nil {
		utils.FastJSON(ctx, 500, gin.H{"message": "取得類別列表失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(ctx, 200, gin.H{"categories": categories})
}

func HandlePostCategory(ctx *gin.Context) {
	var category obj.Category
	if err := utils.FastShouldBindJSON(ctx, &category); err != nil {
		utils.FastJSON(ctx, 400, gin.H{"message": "新增類別失敗\n" + err.Error()})
		return
	}
	if err := db.AddCategory(&category); err != nil {
		utils.FastJSON(ctx, 500, gin.H{"message": "新增類別失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(ctx, 200, gin.H{"message": "新增類別成功"})
}

func HandlePutCategory(ctx *gin.Context) {
	id := uint(conv.ParseInt(ctx.Param("id")))
	if id == 0 {
		utils.FastJSON(ctx, 400, gin.H{"message": "更新類別失敗\nID錯誤"})
		return
	}
	var category obj.Category
	if err := utils.FastShouldBindJSON(ctx, &category); err != nil {
		utils.FastJSON(ctx, 400, gin.H{"message": "更新類別失敗\n" + err.Error()})
		return
	}
	category.ID = id
	if err := db.UpdateCategory(&category); err != nil {
		utils.FastJSON(ctx, 500, gin.H{"message": "更新類別失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(ctx, 200, gin.H{"message": "更新類別成功"})
}

func HandleDeleteCategory(ctx *gin.Context) {
	id := uint(conv.ParseInt(ctx.Param("id")))
	if id == 0 {
		utils.FastJSON(ctx, 400, gin.H{"message": "刪除類別失敗\nID錯誤"})
		return
	}
	category := obj.Category{ID: id}
	if err := db.DeleteCategory(&category); err != nil {
		utils.FastJSON(ctx, 500, gin.H{"message": "刪除類別失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(ctx, 200, gin.H{"message": "刪除類別成功"})
}
