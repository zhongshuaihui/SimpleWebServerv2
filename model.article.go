package main

import "errors"

type article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// v1: the airticle store in disk
var articleList = []article{
	{Id: 1, Title: "Article 1", Content: "The first article"},
	{Id: 2, Title: "Article 2", Content: "The second article"},
}

func getAllArticles() []article {
	return articleList
}

func findArticleById(id int) (article_add *article, err error) {
	for _, a := range articleList {
		if a.Id == id {
			article_add = &a
			err = nil
			return
		}
	}
	return nil, errors.New("article not found")
}

func appendArticle(title string, content string) error {
	a := article{Title: title, Content: content}

	articleList = append(articleList, a)
	return nil
}
