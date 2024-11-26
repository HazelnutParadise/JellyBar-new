package router

import (
	"embed"
	"encoding/base64"
	"io/fs"
	"log"
	"net/http"

	"jellybar/build"
	"jellybar/db"

	"github.com/gin-gonic/gin"
	"github.com/nichady/golte"
)

func GinRouter(siteName string, assetsDir embed.FS) http.Handler {
	// Gin doesn't have a function to wrap middleware, so define our own
	wrapMiddleware := func(middleware func(http.Handler) http.Handler) func(ctx *gin.Context) {
		return func(ctx *gin.Context) {
			middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx.Request = r
				ctx.Next()
			})).ServeHTTP(ctx.Writer, ctx.Request)
			if golte.GetRenderContext(ctx.Request) == nil {
				ctx.Abort()
			}
		}
	}

	// since gin doesm't use stdlib-compatible signatures, we have to wrap them
	// page := func(c string) gin.HandlerFunc {
	// 	return gin.WrapH(golte.Page(c))
	// }
	// layout := func(c string) gin.HandlerFunc {
	// 	return wrapMiddleware(golte.Layout(c))
	// }

	// 使用一個新的變量來存儲 fs.Sub 的結果
	subAssetsDir, err := fs.Sub(assetsDir, "src/assets")
	if err != nil {
		log.Fatalf("Failed to get assets directory: %v", err)
	}

	// 嘗試讀取 logo，如果失敗則使用 nil
	var logo []byte
	if logoBytes, err := fs.ReadFile(subAssetsDir, "logo.png"); err == nil {
		logo = logoBytes
	} else {
		log.Printf("Warning: Could not read logo file: %v", err)
	}

	r := gin.Default()
	// register the main Golte middleware
	r.Use(wrapMiddleware(build.Golte))

	r.Use(checkDBConnection(siteName, logo))

	// 使用 subAssetsDir 而不是 assetsDir
	defineRoutes(r, siteName, subAssetsDir)

	return r
}

func checkDBConnection(siteName string, logo []byte) gin.HandlerFunc {
	if db.IsDBConnected() {
		return nil
	}
	return func(ctx *gin.Context) {
		data := map[string]any{
			"siteName":            siteName,
			"announcementTitle":   "無法連接資料庫",
			"announcementContent": "請稍後再試，或聯繫管理員。",
			"announcementType":    "error",
		}

		// 只有在 logo 存在時才加入 base64 編碼的圖片
		if logo != nil {
			data["siteLogo_base64"] = base64.StdEncoding.EncodeToString(logo)
		}

		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Announcement", data)
		ctx.Abort()
	}
}
