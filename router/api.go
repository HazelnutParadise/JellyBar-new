package router

import (
	"jellybar/apiHandler"

	"github.com/gin-gonic/gin"
)

func defineApi(r *gin.RouterGroup) {
	r.GET("/categories", func(c *gin.Context) {
		apiHandler.HandleGetCategories(c, false)
	})
	r.GET("/article", func(c *gin.Context) {
		apiHandler.HandleGetArticles(c, true)
	})

	adminApi := r.Group("/admin")
	defineAdminApi(adminApi)
}

func defineAdminApi(r *gin.RouterGroup) {
	r.GET("/users", apiHandler.HandleGetUserList)
	r.POST("/user", apiHandler.HandlePostUser)
	r.DELETE("/user/:id", apiHandler.HandleDeleteUser)
	r.PUT("/user/:id", apiHandler.HandleUpdateUser)
	r.POST("/category", apiHandler.HandlePostCategory)
	r.PUT("/category/:id", apiHandler.HandlePutCategory)
	r.DELETE("/category/:id", apiHandler.HandleDeleteCategory)
	r.GET("/article", func(c *gin.Context) {
		apiHandler.HandleGetArticles(c, false)
	})
	r.POST("/article", apiHandler.HandleAddArticle)
}
