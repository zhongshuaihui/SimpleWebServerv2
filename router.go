package main

import (
	"simplewebserverv2/handles"
	"simplewebserverv2/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func initRoutes() {

	router.Use(middleware.SetUserState)

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{"title": "Home Page"},
	// 	)
	// })

	router.GET("/", handles.ShowIndexPage)

	// router.GET("/article/view/:article_id", showArticleDetail)

	// Group article related routes together
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn, handles.ShowLoginPage)
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn, handles.Login)
		userRoutes.GET("/logout", middleware.EnsureLoggedIn, handles.Logout)
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn, handles.ShowRegisterPage)
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn, handles.Register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", handles.ShowArticleDetail)
		articleRoutes.GET("/create", middleware.EnsureLoggedIn, handles.ShowCreateArticlePage)
		articleRoutes.POST("/create", middleware.EnsureLoggedIn, handles.PublishArticle)
	}
}
