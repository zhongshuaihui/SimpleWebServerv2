package handles

import (
	"log"
	"net/http"
	"strconv"

	"simplewebserverv2/middleware"
	"simplewebserverv2/model"

	"github.com/gin-gonic/gin"
)

// show main page
func ShowIndexPage(c *gin.Context) {
	articles, err := model.GetAllArticles()
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
	middleware.Render(c, gin.H{"title": "Home Page", "payload": articles}, "index.html")
}

// show article detail page
func ShowArticleDetail(c *gin.Context) {
	id_str := c.Param("article_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// fmt.Println(id)
		article, err := model.FindArticleById(id)
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
			middleware.Render(c, gin.H{"title": article.Title, "payload": article}, "article.html")
		}
	}

}

func ShowCreateArticlePage(c *gin.Context) {
	middleware.Render(c, gin.H{"title": "Create new article"}, "create_article.html")
}

func PublishArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if err := model.AppendArticle(title, content); err == nil {
		middleware.Render(c, gin.H{"title": "Successfully submit article"}, "submission_successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
