package router

import (
	"jellybar/apiHandler"

	"github.com/gin-gonic/gin"
)

func defineApi(r *gin.RouterGroup) {
	r.GET("/categories", apiHandler.HandleGetCategories)
	adminApi := r.Group("/admin")
	defineAdminApi(adminApi)
}

func defineAdminApi(r *gin.RouterGroup) {
	r.GET("/users", apiHandler.HandleGetUserList)
	r.POST("/user", apiHandler.HandlePostUser)
	r.DELETE("/user/:id", apiHandler.HandleDeleteUser)
	r.PUT("/user/:id", apiHandler.HandleUpdateUser)
	r.POST("/category", apiHandler.HandlePostCategory)
	r.DELETE("/category/:id", apiHandler.HandleDeleteCategory)
}
