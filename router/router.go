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

func GinRouter(siteName string, assetsDir embed.FS, mode int) http.Handler {
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

	// 設定404頁面
	r.NoRoute(handle404(siteName, logo))

	// 使用 subAssetsDir 而不是 assetsDir
	defineRoutes(r, siteName, subAssetsDir, mode)

	return r
}

func checkDBConnection(siteName string, logo []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if db.IsDBConnected() {
			// 數據庫已連接，繼續下一個處理器
			ctx.Next()
			return
		}

		// 數據庫未連接，顯示錯誤頁面
		data := map[string]any{
			"siteName":            siteName,
			"announcementTitle":   "無法連接資料庫",
			"announcementContent": "網站可能還在初始化，請稍後再試。若持續發生，請聯繫管理員。",
			"announcementType":    "error",
		}

		if logo != nil {
			data["siteLogo_base64"] = base64.StdEncoding.EncodeToString(logo)
		}

		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Announcement", data)
		ctx.Abort()
	}
}

func handle404(siteName string, logo []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := map[string]any{
			"siteName":            siteName,
			"announcementTitle":   "404 Not Found 頁面不存在",
			"announcementContent": "怎麼了？你消失了？！我們找不到您試圖前往的頁面，請檢查網址是否正確。",
			"announcementType":    "info",
		}
		if logo != nil {
			data["siteLogo_base64"] = base64.StdEncoding.EncodeToString(logo)
		}
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Announcement", data)
	}
}

func alertDevMode(mode int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果非ajax，則顯示開發模式警告
		if mode == db.DEV && ctx.Request.Header.Get("X-Requested-With") != "XMLHttpRequest" {
			ctx.Writer.Write([]byte(`
		<div style='position: fixed; bottom: 0; left: 0; width: 100%; height: 15px; text-align: center; background-color: yellow; z-index: 9999;'>
		未偵測到環境變數，正在使用開發模式
		</div>`))
		}
		ctx.Next()
	}
}
