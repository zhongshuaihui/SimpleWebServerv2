package main

func initRoutes() {
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{"title": "Home Page"},
	// 	)
	// })

	router.GET("/", showIndexPage)
}
