package router

import (
	"fmt"
	"jellybar/db"
	"jellybar/obj"
	"net/http"
	"strconv"
	"time"

	"github.com/HazelnutParadise/sveltigo"
	"github.com/gin-gonic/gin"
	// sveltigo "github.com/nichady/golte"
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

		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/AdminEditArticle", map[string]any{
			"siteName":    siteName,
			"title":       "編輯文章",
			"thisArticle": thisArticle,
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
}
