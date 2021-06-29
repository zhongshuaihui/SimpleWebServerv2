package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserState() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggendInInterface, _ := c.Get("is_logged_in")
		loggendin := loggendInInterface.(bool)
		if !loggendin {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggendInInterface, _ := c.Get("is_logged_in")
		loggendin := loggendInInterface.(bool)
		if loggendin {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
