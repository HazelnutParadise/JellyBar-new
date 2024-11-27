package router

import (
	"jellybar/app"

	"github.com/gin-gonic/gin"
)

func defineApi(r *gin.RouterGroup) {
	r.GET("/users", app.HandleGetUserList)
	r.POST("/user", app.HandlePostUser)
}
