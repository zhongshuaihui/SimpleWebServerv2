package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	initRoutes()

	router.Run(":8080")
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// respond with json
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// respond with xml
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func render_bad(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// respond with json
		c.JSON(http.StatusBadRequest, data["payload"])
	case "application/xml":
		// respond with xml
		c.XML(http.StatusBadRequest, data["payload"])
	default:
		c.HTML(http.StatusBadRequest, templateName, data)
	}
}
