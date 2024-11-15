package main

import (
	"embed"
	"fmt"
	"jellybar/router"
	"net/http"
)

//go:embed src/assets/*
var assetsDir embed.FS

const siteName = "繽果吉樂 BAR"

func main() {
	r := router.GinRouter(siteName, assetsDir)

	fmt.Println("Serving on :8000")
	http.ListenAndServe(":8000", r)
}
