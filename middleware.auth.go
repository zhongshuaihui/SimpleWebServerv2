package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setUserState(c *gin.Context) {
	token, err := c.Cookie("token")
	if err == nil || token != "" {
		c.Set("is_logged_in", true)
	} else {
		c.Set("is_logged_in", false)
	}
}

func ensureLoggedIn(c *gin.Context) {
	loggendInInterface, _ := c.Get("is_logged_in")
	loggendin := loggendInInterface.(bool)
	if !loggendin {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func ensureNotLoggedIn(c *gin.Context) {
	loggendInInterface, _ := c.Get("is_logged_in")
	loggendin := loggendInInterface.(bool)
	if loggendin {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
