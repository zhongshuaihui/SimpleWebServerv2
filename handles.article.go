package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// show main page
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	// c.HTML(
	// 	http.StatusOK,
	// 	"index.html",
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )

	// use render to handle different request
	render(c, gin.H{"payload": articles}, "index.html")
}

// show article detail page
func showArticleDetail(c *gin.Context) {
	id_str := c.Param("article_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// fmt.Println(id)
		article, err := findArticleById(id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			// c.HTML(
			// 	http.StatusOK,
			// 	"article.html",
			// 	gin.H{
			// 		"payload": article,
			// 	},
			// )

			// use render to handle different request
			render(c, gin.H{"payload": article}, "article.html")
		}
	}

}

func showCreateArticlePage(c *gin.Context) {

}

func appendArticle(c *gin.Context) {

}
