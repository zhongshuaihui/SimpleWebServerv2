package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/webserver?parseTime=true")
	fmt.Println(db, "---------")
	if err != nil {
		log.Fatal(err.Error()) //输出错误，Fatal和panic的区别在前者不会执行defer
	}
	// defer db.Close()
	err = db.Ping()
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
