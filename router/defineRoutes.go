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
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType": "articles",
				"title":    "文章列表",
			},
		})
	})

	r.GET("/article/:id", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Article", map[string]any{
			"siteName": siteName,
			"article": obj.Article{
				Title:       "testjkjkknknknknkmnmknkmnmnjmnmnmnmnmnmnmnjbhjbjefbcjebnfjebnfjewnejcejcnjencnjencjencejjncnecejcnejnncejncjncjenej",
				Description: "testDescriptionlłlll	lłll	llllllllllllllknkjnknknknknknknknknkn",
				PublishDate: time.Now().Format(time.DateOnly),
				UpdateDate:  time.Now().Format(time.DateOnly),
				Content:     "<h1>testContent</h1>",
				Url:         "test",
				ButtonText:  "test",
			},
			"category": obj.Category{
				Title: "testCategory",
			},
			"author": obj.Author{
				Name: "testAuthor",
			},
			"categories": []obj.Category{
				{
					Title: "testCategory",
				},
			},
			"latestArticles": []obj.Article{
				{
					Title: "testLatestArticle",
				},
			},
		})
	})

	r.GET("/category", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType":    "categories",
				"title":       "類別列表",
				"description": "test",
				"items": []obj.Category{
					{
						Title:       "test",
						Description: "test",
						Url:         "test",
						ButtonText:  "test",
					},
					{
						Title:       "test2",
						Description: "test2",
						Url:         "test2",
						ButtonText:  "test",
					},
				},
				"htmlContent": `<h1>test</h1>
			<p>test</p>
			<li>test</li>
			<style>
				h1 {
					color: red;
				}
			</style>
			`,
			},
		})
	})

	r.GET("/category/:id", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType": "category",
				"title":    "文章分類",
				"items": []obj.Article{
					{
						Title: "test",
					},
				},
			},
		})
	})

	r.GET("/author", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType": "author",
				"title":    "作者帳號",
				"items": []obj.Author{
					{
						Name:        "testAuthor",
						Description: "testDescription",
					},
				},
			},
		})
	})
}
