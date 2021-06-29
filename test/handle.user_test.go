package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"simplewebserverv2/handles"
	"simplewebserverv2/middleware"
	"strconv"
	"strings"
	"testing"
)

func TestShowLoginPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/u/login", middleware.EnsureNotLoggedIn(), handles.ShowLoginPage)
	req, _ := http.NewRequest("GET", "/u/login", nil)
	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Login"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowLoginPageAuthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/u/login", middleware.EnsureNotLoggedIn(), handles.ShowLoginPage)
	req, _ := http.NewRequest("GET", "/u/login", nil)
	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 401
		return w.Code == http.StatusUnauthorized
	})
}

func TestLoginUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/u/login", middleware.EnsureNotLoggedIn(), handles.Login)

	// create a request to send the above route
	loginPayload := getLoginPostPayLoad()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))
	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Successfully login"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Successfully login</title>") > 0

		return statusOK && pageOK
	})
}

func TestLoginAuthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/u/login", middleware.EnsureNotLoggedIn(), handles.Login)

	// create a request to send the above route
	loginPayload := getLoginPostPayLoad()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))
	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 401
		return w.Code == http.StatusUnauthorized
	})
}

// test the login function with incorrect username & password
func TestLoginIncorrectUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/u/login", middleware.EnsureNotLoggedIn(), handles.Login)

	// create a request to send the above route
	loginPayload := getRegisterPostPayLoad()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))
	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		return w.Code == http.StatusBadRequest
	})
}

func TestLoginIncorrectAuthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/u/login", middleware.EnsureNotLoggedIn(), handles.Login)

	// create a request to send the above route
	loginPayload := getRegisterPostPayLoad()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))
	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 401
		return w.Code == http.StatusUnauthorized
	})
}

func TestShowRegisterPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/u/register", middleware.EnsureNotLoggedIn(), handles.ShowRegisterPage)
	req, _ := http.NewRequest("GET", "/u/register", nil)
	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Register"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})
}

func TestShowRegisterPageAuthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/u/register", middleware.EnsureNotLoggedIn(), handles.ShowRegisterPage)
	req, _ := http.NewRequest("GET", "/u/register", nil)
	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 401
		return w.Code == http.StatusUnauthorized
	})
}

func TestRegisterUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/u/register", middleware.EnsureNotLoggedIn(), handles.Register)

	// create a request to send the above route
	registerPayload := getRegisterPostPayLoad()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registerPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registerPayload)))
	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		statusOK := w.Code == http.StatusOK

		// Test whether the page title is "Successfully register & login"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Successfully register and login</title>") > 0

		return statusOK && pageOK
	})
}

func TestRegisterAuthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/u/register", middleware.EnsureNotLoggedIn(), handles.Login)

	// create a request to send the above route
	registerPayload := getRegisterPostPayLoad()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registerPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registerPayload)))
	testHTTPResponseAuthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 401
		return w.Code == http.StatusUnauthorized
	})
}

// test the register function with incorrect username & password
func TestRegisterIncorrectUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/u/register", middleware.EnsureNotLoggedIn(), handles.Register)

	// create a request to send the above route
	registerPayload := getLoginPostPayLoad()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registerPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registerPayload)))
	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test whether the http state code is 200
		return w.Code == http.StatusBadRequest
	})
}

func getLoginPostPayLoad() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "123")
	return params.Encode()
}

func getRegisterPostPayLoad() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "123")
	return params.Encode()
}
