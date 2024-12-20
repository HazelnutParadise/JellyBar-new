package main

import (
	"embed"
	"fmt"
	"jellybar/db"
	"jellybar/router"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed src/assets/*
var assetsDir embed.FS

const siteName = "繽果吉樂 BAR"

func main() {
	modeEnv := os.Getenv("MODE")
	mode := db.DEV
	if modeEnv == "PROD" {
		mode = db.PROD
		gin.SetMode(gin.ReleaseMode)
	}
	go db.InitDB(mode)
	r := router.GinRouter(siteName, &assetsDir, mode)

	fmt.Println("Serving on :3331")
	http.ListenAndServe(":3331", r)
}
