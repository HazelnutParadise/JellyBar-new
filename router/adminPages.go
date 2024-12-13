package router

import (
	"fmt"
	"jellybar/db"
	"net/http"
	"strconv"

	"github.com/HazelnutParadise/sveltigo"
	"github.com/gin-gonic/gin"
)

func defineAdminPages(r *gin.RouterGroup, siteName string) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/admin/articles")
	})

	r.GET("/articles", func(ctx *gin.Context) {
		categories, err := db.GetCategories(db.GetCategoriesOption{
			PreloadArticles: true,
			OnlyPublished:   false,
		})
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/AdminArticleAndCategory", map[string]any{
			"siteName":       siteName,
			"title":          "文章與類別",
			"categoriesData": categories,
		})
	})

	r.GET("/article/new", func(ctx *gin.Context) {
		categories, err := db.GetCategories(db.GetCategoriesOption{
			PreloadArticles: false,
		})
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/AdminEditArticle", map[string]any{
			"siteName":   siteName,
			"title":      "新增文章",
			"categories": categories,
			"isNew":      true,
		})
	})

	r.GET("/article/edit/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		article, err := db.GetArticleByID(uint(idInt), false)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if article == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		categories, err := db.GetCategories(db.GetCategoriesOption{
			PreloadArticles: false,
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/AdminEditArticle", map[string]any{
			"siteName":    siteName,
			"title":       "編輯文章",
			"thisArticle": article,
			"categories":  categories,
			"isNew":       false,
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		users, err := db.GetUsers()
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/AdminUser", map[string]any{
			"siteName": siteName,
			"title":    "用戶管理",
			"users":    users,
		})
	})

	r.GET("/my-author-page", func(ctx *gin.Context) {
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/MyAuthorPage", map[string]any{
			// "siteName": siteName,
			// "title":    "作者資料",
		})
	})
}
