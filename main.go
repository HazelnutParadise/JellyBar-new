package main

import (
	"embed"
	"fmt"
	"jellybar/router"
	"net/http"
)

//go:embed src/assets/*
var assetsDir embed.FS

func main() {
	r := router.GinRouter(assetsDir)

	fmt.Println("Serving on :8000")
	http.ListenAndServe(":8000", r)
}
