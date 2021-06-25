package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var Db *sql.DB

func main() {
	var err error
	// v2 read data from database
	Db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/webserver?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer Db.Close()
	err = Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

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
