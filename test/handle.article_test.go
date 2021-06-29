package tests

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"simplewebserverv2/handles"
	"simplewebserverv2/middleware"
	"simplewebserverv2/model"
	"strconv"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handles.ShowIndexPage)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowIndexPageAuthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handles.ShowIndexPage)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", handles.ShowArticleDetail)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Article 1"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleAuthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", handles.ShowArticleDetail)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Article 1"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}

// test a Get request to the home page returns the list of articles in json format
// when the accept header is set to application/json
func TestShowIndexPageJson(t *testing.T) {
	r := getRouter(true)

	r.GET("/", handles.ShowIndexPage)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test the response of json
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []model.Article
		err = json.Unmarshal(p, &articles)

		return statusOK && err == nil && len(articles) >= 2
	})
}

func TestArticleXml(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", handles.ShowArticleDetail)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test the response of json
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var article model.Article
		err = xml.Unmarshal(p, &article)

		return statusOK && err == nil && len(article.Title) >= 0 && article.Id == 1
	})
}

func TestCreateArticlePageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/create", middleware.EnsureLoggedIn(), handles.ShowCreateArticlePage)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/article/create", nil)

	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 401
		return w.Code == http.StatusUnauthorized
	})
}

func TestCreateArticlePageAuthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/create", middleware.EnsureLoggedIn(), handles.ShowCreateArticlePage)

	// create a request to send above route
	req, _ := http.NewRequest("GET", "/article/create", nil)

	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Create new article"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Create new article</title>") > 0

		return statusOK && pageOK
	})
}

func TestPublishArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.POST("/article/create", middleware.EnsureLoggedIn(), handles.PublishArticle)

	// create a request to send the above route
	articlePayload := getArticlePostPayLoad()
	req, _ := http.NewRequest("POST", "/article/create", strings.NewReader(articlePayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(articlePayload)))

	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 401
		return w.Code == http.StatusUnauthorized
	})
}

func TestPublishArticleAuthenticated(t *testing.T) {
	r := getRouter(true)

	r.POST("/article/create", middleware.EnsureLoggedIn(), handles.PublishArticle)

	// create a request to send the above route
	articlePayload := getArticlePostPayLoad()
	req, _ := http.NewRequest("POST", "/article/create", strings.NewReader(articlePayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(articlePayload)))

	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Successfully submit article"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Successfully submit article</title>") > 0

		return statusOK && pageOK
	})
}

func getArticlePostPayLoad() string {
	params := url.Values{}
	params.Add("title", "test article title")
	params.Add("content", "test article content")
	return params.Encode()
}
