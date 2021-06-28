package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserState(c *gin.Context) {
	token, err := c.Cookie("token")
	if err == nil || token != "" {
		c.Set("is_logged_in", true)
	} else {
		c.Set("is_logged_in", false)
	}
}

func EnsureLoggedIn(c *gin.Context) {
	loggendInInterface, _ := c.Get("is_logged_in")
	loggendin := loggendInInterface.(bool)
	if !loggendin {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func EnsureNotLoggedIn(c *gin.Context) {
	loggendInInterface, _ := c.Get("is_logged_in")
	loggendin := loggendInInterface.(bool)
	if loggendin {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
