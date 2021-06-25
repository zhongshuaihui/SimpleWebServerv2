package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {
	render(c, gin.H{}, "login.html")
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isUserValid(username, password) {
		// generate the random token, this method is not safe
		token := strconv.FormatInt(rand.Int63(), 16)

		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{}, "login_successful.html")
	} else {
		render_bad(c, gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "No such user or password error",
		}, "login.html")
	}
}

func logout(c *gin.Context) {
	// clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
