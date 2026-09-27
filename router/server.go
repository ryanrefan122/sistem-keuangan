package router

import (
	"fmt"
	"main/handler"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	// Public
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.GET("/artikel", handler.GetArticle)

	// Protected
	user := router.Group("/")
	user.Use(middleware.AuthMiddleware())

	{
		user.GET("/profil", handler.ProfilUser)
		user.PUT("/profil", handler.UpdateUser)
		user.DELETE("/profil", handler.DeleteUser)

		user.POST("/artikel", handler.CreateArticle)
		user.GET("/my/artikel", handler.GetArticleUser)

		user.GET("/artikel/:id", handler.GetArticleById)
		user.PUT("/my/artikel/:id", handler.UpdatedArticle)
		user.DELETE("/my/artikel/:id", handler.DeleteArticle)
	}
	fmt.Println("berhasil tersambung")

	router.Run(":8080")

}
