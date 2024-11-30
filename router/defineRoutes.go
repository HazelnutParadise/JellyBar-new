package router

import (
	"io/fs"
	"net/http"
	"time"

	"jellybar/md"
	"jellybar/obj"

	"github.com/gin-gonic/gin"
	"github.com/nichady/golte"
)

func defineRoutes(r *gin.Engine, siteName string, assets *fs.FS, logoBase64 *string) {
	r.StaticFS("/assets", http.FS(*assets))

	pages := r.Group("/")
	pages.Use(func(ctx *gin.Context) {
		golte.AddLayout(ctx.Request, "layouts/NavbarAndFooter", map[string]any{
			"siteName": siteName,
			"logo":     logoBase64,
		})
	})
	pages.GET("/", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Index", map[string]any{
			"siteName": siteName,
			"title":    "歡迎光臨 Jelly Bar",
		})
	})

	pages.GET("/themes", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Themes", map[string]any{
			"siteName": siteName,
			"title":    "主題展示",
		})
	})

	pages.GET("/article", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType": "articles",
				"title":    "文章列表",
				"items": []obj.ListItem{
					{
						Title:       "test",
						Description: "test",
						Icon:        "test",
						Url:         "test",
						ButtonText:  "test",
					},
				},
			},
		})
	})

	pages.GET("/article/:id", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Article", map[string]any{
			"siteName": siteName,
			"article": obj.Article{
				Title:       "testjkjkknknknknkmnmknkmnmnjmnmnmnmnmnmnmnjbhjbjefbcjebnfjebnfjewnejcejcnjencnjencjencejjncnecejcnejnncejncjncjenej",
				Description: "testDescriptionlłlll	lłll	llllllllllllllknkjnknknknknknknknknknkn",
				PublishAt:   time.Now(),
				UpdateAt:    time.Now(),
				Content:     md.Parse("<h1>test</h1>\n<p>test</p>\n# markdown h1\n## markdown h2\n### markdown h3\n#### markdown h4\n##### markdown h5\n###### markdown h6\n```python\nprint(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")print(\"test\")\n```\n<li>test</li>\n<style>\n\th1 {\n\t\tcolor: red;\n\t}\n</style>"),
			},
			"category": obj.Category{
				Name: "testCategory",
			},
			"author": obj.Author{
				Name: "testAuthor",
			},
			"categories": []obj.Category{
				{
					Name: "testCategory",
				},
			},
			"latestArticles": []obj.Article{
				{
					Title: "testLatestArticle",
				},
			},
		})
	})

	pages.GET("/category", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType":    "categories",
				"title":       "類別列表",
				"description": "test",
				"items": []obj.ListItem{
					{
						Title:       "test",
						Description: "test",
						Url:         "test",
						ButtonText:  "test",
						Icon:        "test",
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

	pages.GET("/category/:id", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType": "category",
				"title":    "文章分類",
				"items": []obj.ListItem{
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

	pages.GET("/author", func(ctx *gin.Context) {
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

	pages.GET("/search/:keyword", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType": "search",
				"title":    "搜尋結果",
				"items": []obj.ListItem{
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

	pages.GET("/login", func(ctx *gin.Context) {
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Login", map[string]any{
			"siteName": siteName,
		})
	})

	admin := r.Group("/admin", func(ctx *gin.Context) {
		golte.AddLayout(ctx.Request, "layouts/AdminNavbarAndFooter", map[string]any{
			"siteName": siteName,
			"logo":     logoBase64,
		})
	})
	defineAdminPages(admin, siteName)

	api := r.Group("/api")
	defineApi(api)

}
