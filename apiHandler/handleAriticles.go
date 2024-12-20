package apiHandler

import (
	"jellybar/db"
	"jellybar/obj"
	"jellybar/utils"

	"github.com/gin-gonic/gin"
)

func HandleGetArticles(c *gin.Context, onlyPublished bool) {
	categoryId := c.Query("categoryId")
	articles, err := db.GetArticles(categoryId, onlyPublished)
	if err != nil {
		utils.FastJSON(c, 500, gin.H{"message": "取得文章列表失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(c, 200, gin.H{"articles": articles})
}

func HandleAddArticle(c *gin.Context) {
	var article obj.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.FastJSON(c, 400, gin.H{"message": "新增文章失敗\n" + err.Error()})
		return
	}
	if err := db.AddArticle(&article); err != nil {
		utils.FastJSON(c, 500, gin.H{"message": "新增文章失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(c, 200, gin.H{"message": "新增文章成功"})
}

func HandleUpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article obj.Article
	if err := utils.FastShouldBindJSON(c, &article); err != nil {
		utils.FastJSON(c, 400, gin.H{"message": "更新文章失敗\n" + err.Error()})
		return
	}
	if err := db.UpdateArticle(id, &article); err != nil {
		utils.FastJSON(c, 500, gin.H{"message": "更新文章失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(c, 200, gin.H{"message": "更新文章成功"})
}

func HandleDeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if err := db.DeleteArticle(id); err != nil {
		utils.FastJSON(c, 500, gin.H{"message": "刪除文章失敗\n" + err.Error()})
		return
	}
	utils.FastJSON(c, 200, gin.H{"message": "刪除文章成功"})
}
