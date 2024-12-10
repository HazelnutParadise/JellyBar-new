package router

import (
	"fmt"
	"io/fs"
	"net/http"

	"jellybar/obj"

	"jellybar/db"

	"github.com/HazelnutParadise/Go-Utils/conv"
	"github.com/HazelnutParadise/sveltigo"
	"github.com/gin-gonic/gin"
)

func defineRoutes(r *gin.Engine, siteName string, assets *fs.FS, logoBase64 *string) {
	r.StaticFS("/assets", http.FS(*assets))

	pages := r.Group("/")
	pages.Use(func(ctx *gin.Context) {
		sveltigo.AddLayout(ctx.Request, "layouts/NavbarAndFooter", map[string]any{
			"siteName": siteName,
			"logo":     logoBase64,
		})
	})
	pages.GET("/", func(ctx *gin.Context) {
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/Index", map[string]any{
			"siteName": siteName,
			"title":    "歡迎光臨 Jelly Bar",
		})
	})

	pages.GET("/themes", func(ctx *gin.Context) {
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/Themes", map[string]any{
			"siteName": siteName,
			"title":    "主題展示",
		})
	})

	pages.GET("/article", func(ctx *gin.Context) {
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
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
		id := uint(conv.ParseInt(ctx.Param("id")))
		article, err := db.GetArticleByID(id, true)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if article == nil {
			ctx.Request.URL.Path = "/not_found"
			r.HandleContext(ctx)
			return
		}
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/Article", map[string]any{
			"siteName": siteName,
			"article":  article,
			"category": article.Category,
			"author": obj.Author{ //TODO
				Name: "testAuthor",
			},
			//sidebar
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
		categories, err := db.GetCategories(db.GetCategoriesOption{
			PreloadArticles: false,
		})
		if err != nil {
			fmt.Println(err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		var items []obj.ListItem
		for i := range *categories {
			category := (*categories)[i]
			items = append(items, obj.ListItem{
				Title:       category.Name,
				Description: category.Description,
				Url:         "/category/" + conv.ToString(category.ID),
				ButtonText:  "查看",
				Icon:        "📚",
			})
		}
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType":    "categories",
				"title":       "文章分類",
				"description": "這裡列出了本站所有文章類別，挑一個來看看吧！",
				"items":       items,
			},
		},
		)
	})

	pages.GET("/category/:id", func(ctx *gin.Context) {
		categoryID := uint(conv.ParseInt(ctx.Param("id")))
		category, err := db.GetCategoryByID(categoryID, true)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if category == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		var items []obj.ListItem
		for i := range category.Articles {
			article := category.Articles[i]
			items = append(items, obj.ListItem{
				Title:       article.Title,
				Description: article.Description,
				Url:         "/article/" + conv.ToString(article.ID),
				ButtonText:  "查看",
				Icon:        "🗒️",
			})
		}
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
			"siteName": siteName,
			"data": map[string]any{
				"pageType":    "category",
				"title":       "[類別] " + category.Name,
				"description": category.Description,
				"items":       items,
			},
		})
	})

	pages.GET("/author", func(ctx *gin.Context) {
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
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
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/PageWithList", map[string]any{
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
		sveltigo.RenderPage(ctx.Writer, ctx.Request, "pages/Login", map[string]any{
			"siteName": siteName,
		})
	})

	admin := r.Group("/admin", func(ctx *gin.Context) {
		sveltigo.AddLayout(ctx.Request, "layouts/AdminNavbarAndFooter", map[string]any{
			"siteName": siteName,
			"logo":     logoBase64,
		})
	})
	defineAdminPages(admin, siteName)

	api := r.Group("/api")
	defineApi(api)

}
