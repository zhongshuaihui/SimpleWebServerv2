package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"simplewebserverv2/middleware"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// set gin to test mode
	gin.SetMode(gin.TestMode)

	middleware.InitDataBase()
	defer middleware.Db.Close()

	// run other tests
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../templates/*")
		r.Use(middleware.SetUserState())
	}
	return r
}

func testHTTPResponseUnauthenticated(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()

	// create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func testHTTPResponseAuthenticated(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()
	// set the token cookie to simulate an authenticated user
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	req.Header.Add("Cookie", "token=123")

	// create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {
	req, _ := http.NewRequest("GET", "/", nil)
	testHTTPResponseUnauthenticated(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return expectedHTTPCode == w.Code
	})
}
