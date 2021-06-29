package handles

import (
	"math/rand"
	"net/http"
	"simplewebserverv2/middleware"
	"simplewebserverv2/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	middleware.Render(c, gin.H{"title": "Login"}, "login.html")
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if model.IsUserValid(username, password) {
		// generate the random token, this method is not safe
		token := strconv.FormatInt(rand.Int63(), 16)

		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		middleware.Render(c, gin.H{"title": "Successfully login"}, "login_successful.html")
	} else {
		middleware.Render_bad(c, gin.H{
			"title":        "Login",
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "No such user or password error",
		}, "login.html")
	}
}

func Logout(c *gin.Context) {
	// clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func ShowRegisterPage(c *gin.Context) {
	middleware.Render(c, gin.H{"title": "Register"}, "register.html")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if err := model.RegisterNewUser(username, password); err == nil {
		token := strconv.FormatInt(rand.Int63(), 16)
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		middleware.Render(c, gin.H{"title": "Successfully register and login"}, "login_successful.html")
	} else {
		middleware.Render_bad(c, gin.H{
			"title":        "Register",
			"ErrorTitle":   "Register Failed",
			"ErrorMessage": err.Error(),
		}, "register.html")
	}
}
