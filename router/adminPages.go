package router

import (
	"net/http"

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
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/AdminEditArticle", map[string]any{
			"siteName": siteName,
			"title":    "編輯文章",
			"id":       id,
		})
	})
}
