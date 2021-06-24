package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "Home Page"},
		)
	})
}
