package router

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"jellybar/md"
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
				"items": []obj.Article{
					{
						Title:       "test",
						Description: "test",
						Url:         "test",
						ButtonText:  "test",
					},
				},
			},
		})
	})

	r.GET("/article/:id", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Article", map[string]any{
			"siteName": siteName,
			"article": obj.Article{
				Title:       "testjkjkknknknknkmnmknkmnmnjmnmnmnmnmnmnmnjbhjbjefbcjebnfjebnfjewnejcejcnjencnjencjencejjncnecejcnejnncejncjncjenej",
				Description: "testDescriptionlłlll	lłll	llllllllllllllknkjnknknknknknknknknknkn",
				PublishDate: time.Now().Format(time.DateOnly),
				UpdateDate:  time.Now().Format(time.DateOnly),
				Content:     md.Parse("<h1>test</h1>\n<p>test</p>\n# markdown h1\n## markdown h2\n### markdown h3\n#### markdown h4\n##### markdown h5\n###### markdown h6\n```python\nprint(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")\n```\n<li>test</li>\n<style>\n\th1 {\n\t\tcolor: red;\n\t}\n</style>"),
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
				"htmlContent": md.Parse("<h1>test</h1>\n<p>test</p>\n# markdown h1\n## markdown h2\n### markdown h3\n#### markdown h4\n##### markdown h5\n###### markdown h6\n```python\nprint(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")\n```\n<li>test</li>\n<style>\n\th1 {\n\t\tcolor: red;\n\t}\n</style>"),
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
						Title:       "test",
						Description: "test",
						Url:         "test",
						ButtonText:  "test",
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

	r.GET("/search/:keyword", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType": "search",
				"title":    "搜尋結果",
				"items": []obj.Article{
					{
						Title:       "test",
						Description: "test",
						Url:         "test",
						ButtonText:  "test",
					},
				},
			},
		})
	})

	admin := r.Group("/admin")
	defineAdminPages(admin, siteName)

}
