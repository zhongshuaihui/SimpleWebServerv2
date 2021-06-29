package tests

import (
	"net/http"
	"net/http/httptest"
	"simplewebserverv2/middleware"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetUserStateUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", middleware.SetUserState(), func(c *gin.Context) {
		// the token is not set, thus the userstate should be false
		loginInterface, exist := c.Get("is_logged_in")
		if exist && loginInterface.(bool) {
			t.Fail()
		}
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
}

func TestSetUserStateAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", middleware.SetUserState(), func(c *gin.Context) {
		// the token is set, thus the userstate should be true
		loginInterface, exist := c.Get("is_logged_in")
		if !exist || !loginInterface.(bool) {
			t.Fail()
		}
	})
	w := httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Cookie", "token=123")
	r.ServeHTTP(w, req)
}

func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), middleware.EnsureLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})
	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func TestEnsureLoggedInAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), middleware.EnsureLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestEnsureNotLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), middleware.EnsureNotLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestEnsureNotLoggedInAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), middleware.EnsureNotLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})
	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}
