package router

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"jellybar/obj"

	"github.com/gin-gonic/gin"
	"github.com/nichady/golte"
)

func defineRoutes(r *gin.Engine, siteName string, assetsDir embed.FS) {
	assets, err := fs.Sub(assetsDir, "src/assets")
	if err != nil {
		fmt.Printf("Error getting sub filesystem: %v\n", err)
	}
	r.StaticFS("/assets", http.FS(assets))

	r.GET("/", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/App", map[string]any{
			"siteName": siteName,
			"title":    "歡迎光臨 Jelly Bar",
		})
	})

	r.GET("/themes", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Themes", map[string]any{
			"siteName": siteName,
			"title":    "主題展示",
		})
	})

	r.GET("/article", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Article", map[string]any{
			"siteName": siteName,
			"article": obj.Article{
				Title:       "test",
				Description: "testDescription",
				PublishDate: time.Now().Format(time.DateOnly),
				Content:     "<h1>testContent</h1>",
				Url:         "test",
				Icon:        "🍋",
				ButtonText:  "test",
			},
			"category": obj.Category{
				Title: "testCategory",
			},
			"author": obj.Author{
				Name: "testAuthor",
			},
		})
	})

	r.GET("/category", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Categories", map[string]any{
			"siteName":    siteName,
			"theme":       "sunset",
			"title":       "文章分類",
			"description": "test",
			"icon":        "🍋",
			"categories": []obj.Category{
				{
					Title:       "test",
					Description: "test",
					Url:         "test",
					Icon:        "🍋",
					ButtonText:  "test",
				},
				{
					Title:       "test2",
					Description: "test2",
					Url:         "test2",
					Icon:        "🍋",
					ButtonText:  "test",
				},
			},
		})
	})

	r.GET("/author", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Author", map[string]any{
			"siteName": siteName,
			"title":    "作者帳號",
		})
	})
}
