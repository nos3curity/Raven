package controllers

import (
	"ac-130/models"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	context "github.com/beego/beego/v2/server/web/context"
)

type LoginController struct {
	beego.Controller
}

func ValidateJwtFilter(c *context.Context) {

	if c.Request.URL.Path == "/login" {
		return
	}

	// Check if the "token" cookie is set
	cookie, err := c.Request.Cookie("session")
	if err != nil || cookie.Value == "" {
		// If the cookie is not set, redirect to the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Call ValidateJwt with the token string
	valid, err := models.ValidateJwt(cookie.Value)
	if err != nil || !valid {
		// If the token is not valid or an error occurred, redirect to the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}

	return
}

func (c *LoginController) Get() {

	c.TplName = "user/login.html"
	return
}

func (c *LoginController) SignIn() {

	// Parse username and password
	username := c.GetString("username")
	password := c.GetString("password")

	if username == "" || password == "" {
		c.Ctx.WriteString("Provide a username and password")
		return
	}

	passwordValid, err := models.CheckPassword(password)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	if passwordValid == false {
		c.Ctx.WriteString("Invalid credentials")
		return
	}

	// Get the JWT secret
	token, err := models.IssueJwt(username)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Set the client cookie
	c.Ctx.SetCookie(
		"session",
		token,
	)

	c.Redirect("/", 302) // CHANGE AS NEEDED
	return
}
