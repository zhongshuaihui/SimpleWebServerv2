package main

func initRoutes() {

	router.Use(setUserState)

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{"title": "Home Page"},
	// 	)
	// })

	router.GET("/", showIndexPage)

	// router.GET("/article/view/:article_id", showArticleDetail)

	// Group article related routes together
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn, showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn, login)
		userRoutes.GET("/logout", ensureLoggedIn, logout)
		userRoutes.GET("/register", ensureNotLoggedIn, showRegisterPage)
		userRoutes.POST("/register", ensureNotLoggedIn, register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", showArticleDetail)
		articleRoutes.GET("/create", ensureLoggedIn, showCreateArticlePage)
		articleRoutes.POST("/create", ensureLoggedIn, publishArticle)
	}
}
