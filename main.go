package main

import (
	_ "github.com/go-sql-driver/mysql"

	"simplewebserverv2/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	middleware.InitDataBase()

	defer middleware.Db.Close()

	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	initRoutes()

	router.Run(":8080")
}
