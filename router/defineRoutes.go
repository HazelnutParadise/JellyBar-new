package router

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

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
}
