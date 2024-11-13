package router

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

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
			"title":    "繽果吉樂 BAR",
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
			"title":    "文章",
		})
	})

	r.GET("/category", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Category", map[string]any{
			"siteName": siteName,
			"theme":    "sunset",
			"title":    "文章分類",
			"category": map[string]any{
				"name":        "test",
				"description": "test",
				"icon":        "🍋",
			},
			"articles": []obj.Article{
				{
					Title:       "test",
					Description: "test",
					Url:         "test",
					Category:    "test",
					Icon:        "🍋",
					ButtonText:  "test",
				},
				{
					Title:       "test2",
					Description: "test2",
					Url:         "test2",
					Category:    "test",
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
