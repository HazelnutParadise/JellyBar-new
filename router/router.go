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

func GinRouter(siteName string, assetsDir *embed.FS, mode int) http.Handler {
	// Gin doesn't have a function to wrap middleware, so define our own
	wrapMiddleware := func(middleware func(http.Handler) http.Handler) func(ctx *gin.Context) {
		return func(ctx *gin.Context) {
			// 將 ctx.Request 和 ctx.Writer 使用指標引用，避免拷貝
			req := ctx.Request
			writer := ctx.Writer

			// 靜態化處理器，直接使用閉包並保持指標傳遞
			staticHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// 使用指標直接操作 ctx 的 Request
				ctx.Request = r
				ctx.Next()
			})

			// 使用中介軟體處理請求
			handler := middleware(staticHandler)
			handler.ServeHTTP(writer, req)

			// 使用指標直接操作請求進行檢查
			if renderCtx := golte.GetRenderContext(req); renderCtx == nil {
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
	subAssetsDir, err := fs.Sub(*assetsDir, "src/assets")
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

	logoBase64 := base64.StdEncoding.EncodeToString(logo)

	r := gin.Default()
	// register the main Golte middleware
	r.Use(wrapMiddleware(build.Golte))

	r.Use(func(ctx *gin.Context) {
		golte.AddLayout(ctx.Request, "App", map[string]any{
			"logo": logoBase64,
		})
	})

	r.Use(alertDevMode(mode))
	r.Use(checkDBConnection(siteName, &logoBase64, mode))

	// 設定404頁面
	r.NoRoute(handle404(siteName, &logoBase64))

	// 使用 subAssetsDir 而不是 assetsDir
	defineRoutes(r, siteName, &subAssetsDir, &logoBase64)

	return r
}

func checkDBConnection(siteName string, logoBase64 *string, mode int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if db.IsDBConnected(mode) {
			ctx.Next()
			return
		}

		// 先設置狀態碼
		ctx.Status(http.StatusInternalServerError)

		data := map[string]any{
			"siteName":            siteName,
			"announcementTitle":   "無法連接資料庫",
			"announcementContent": "網站可能還在初始化，請稍後再試。若持續發生，請聯繫管理員。",
			"announcementType":    "error",
		}

		if logoBase64 != nil {
			data["siteLogo_base64"] = *logoBase64
		}

		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Announcement", data)
		ctx.Abort()
	}
}

func handle404(siteName string, logoBase64 *string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := map[string]any{
			"siteName":            siteName,
			"announcementTitle":   "404 Not Found 頁面不存在",
			"announcementContent": "怎麼了？你消失了？！我們找不到您試圖前往的頁面，請檢查網址是否正確。",
			"announcementType":    "info",
		}
		if logoBase64 != nil {
			data["siteLogo_base64"] = *logoBase64
		}
		golte.RenderPage(ctx.Writer, ctx.Request, "pages/Announcement", data)
	}
}

func alertDevMode(mode int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果非ajax，則顯示開發模式警告
		if mode == db.DEV {
			golte.AddLayout(ctx.Request, "layouts/DevModeAlert", nil)
		}
		ctx.Next()
	}
}
