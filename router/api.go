package router

import (
	"jellybar/app"

	"github.com/gin-gonic/gin"
)

func defineApi(r *gin.RouterGroup) {
	r.GET("/categories", app.HandleGetCategories)
	adminApi := r.Group("/admin")
	defineAdminApi(adminApi)
}

func defineAdminApi(r *gin.RouterGroup) {
	r.GET("/users", app.HandleGetUserList)
	r.POST("/user", app.HandlePostUser)
	r.DELETE("/user", app.HandleDeleteUser)
	r.PUT("/user", app.HandleUpdateUser)
}
