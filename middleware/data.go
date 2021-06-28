package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDataBase() {
	var err error
	// v2 read data from database
	Db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/webserver?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Render(c *gin.Context, data gin.H, templateName string) {
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

func Render_bad(c *gin.Context, data gin.H, templateName string) {
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
