package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// show main page
func showIndexPage(c *gin.Context) {
	articles, err := getAllArticles()
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(
	// 	http.StatusOK,
	// 	"index.html",
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )

	// use render to handle different request
	// the payload can carry all the articles
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
	render(c, gin.H{"title": "Create new article"}, "create_article.html")
}

func publishArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if err := appendArticle(title, content); err == nil {
		render(c, gin.H{"title": "Successfully submit article"}, "submission_successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
