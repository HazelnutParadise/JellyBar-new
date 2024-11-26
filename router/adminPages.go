package router

import (
	"jellybar/obj"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nichady/golte"
)

func defineAdminPages(r *gin.RouterGroup, siteName string) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/admin/articles")
	})

	r.GET("/articles", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/AdminArticleAndCategory", map[string]any{
			"siteName": siteName,
			"title":    "文章與類別",
		})
	})

	r.GET("/article/new", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/AdminEditArticle", map[string]any{
			"siteName": siteName,
			"title":    "新增文章",
		})
	})

	r.GET("/article/edit/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		thisArticle := obj.Article{
			ID:          uint(idInt),
			Title:       "Title",
			Content:     "Content",
			Description: "Description",
			Status:      "draft",
			PublishAt:   time.Now(),
			UpdateAt:    time.Now(),
			Category: obj.Category{
				ID:   0,
				Name: "Category",
			},
			Media: []string{},
		}

		golte.RenderPage(ctx.Writer, ctx.Request, "pages/AdminEditArticle", map[string]any{
			"siteName":    siteName,
			"title":       "編輯文章",
			"thisArticle": thisArticle,
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/AdminUser", map[string]any{
			"siteName": siteName,
			"title":    "用戶管理",
		})
	})
}
