package main

type article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// v2: the airticle store in disk
var articleList = []article{
	{Id: 1, Title: "Article 1", Content: "The first article"},
	{Id: 2, Title: "Article 2", Content: "The second article"},
}

func getAllArticles() []article {
	return articleList
}
