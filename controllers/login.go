package controllers

import (
	"fmt"
	"net/http"
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
	context "github.com/beego/beego/v2/server/web/context"
)

type LoginController struct {
	beego.Controller
}

func ValidateJwtFilter(c *context.Context) {
	var token string
	var err error

	if c.Request.URL.Path == "/login" {
		return
	}

	// First, try to get the token from the Authorization header
	token = c.Request.Header.Get("Authorization")

	// If the token is not in the Authorization header, check the cookie
	if token == "" {
		cookie, err := c.Request.Cookie("session")
		if err != nil || cookie.Value == "" {
			// If neither Authorization header nor cookie is set, redirect to the login page
			c.Redirect(http.StatusFound, "/login")
			return
		}
		token = cookie.Value
	}

	// Call ValidateJwt with the token string
	valid, err := models.ValidateJwt(token)
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

func (c *LoginController) Profile() {

	// Get all teams
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	// Get the JWT cookie
	jwt, err := c.Ctx.Request.Cookie("session")
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get the username
	username, err := models.GetCurrentUsername(jwt.Value)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	fmt.Println(jwt.Value)

	c.Data["authorization"] = jwt.Value
	c.Data["username"] = username
	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
	c.TplName = "user/profile.html"
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
