package model

import (
	"errors"
	"log"
	"simplewebserverv2/middleware"
)

type article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// v1: the airticle store in disk
// var articleList = []article{
// 	{Id: 1, Title: "Article 1", Content: "The first article"},
// 	{Id: 2, Title: "Article 2", Content: "The second article"},
// }

// v2: the airticle store in database
var articleList []article

func GetAllArticles() ([]article, error) {
	if len(articleList) == 0 {
		rows, err := middleware.Db.Query("select * from article")
		if err != nil {
			return articleList, err
		}
		defer rows.Close()

		for rows.Next() {
			var a article
			rows.Scan(&a.Id, &a.Title, &a.Content)
			articleList = append(articleList, a)
		}
	}

	return articleList, nil
}

func FindArticleById(id int) (article_add *article, err error) {
	for _, a := range articleList {
		if a.Id == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}

func AppendArticle(title string, content string) error {
	stmt, err := middleware.Db.Prepare("insert into article (title, content) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(title, content)
	if err != nil {
		return err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	Id := int(id)

	a := article{Id: Id, Title: title, Content: content}
	articleList = append(articleList, a)

	return nil
}
