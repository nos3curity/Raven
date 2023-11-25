package controllers

import (
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type CommentsController struct {
	beego.Controller
}


func (c *CommentsController) Add() {
    // Extract system IP and comment text from the request
    systemIp := c.GetString("system_ip")
    commentText := c.GetString("comment")

    // Get the JWT cookie
    jwtCookie, err := c.Ctx.Request.Cookie("session")
    if err != nil {
        // Handle error, such as redirecting to login
        c.Ctx.WriteString("Authentication error: " + err.Error())
        return
    }

    // Get the username from the JWT token
    username, err := models.GetCurrentUsername(jwtCookie.Value)
    if err != nil {
        // Handle error
        c.Ctx.WriteString("Error extracting username: " + err.Error())
        return
    }

    // Call the function to add a comment
    err = models.AddComment(systemIp, username, commentText)
    if err != nil {
        // Handle error
        c.Ctx.WriteString("Error adding comment: " + err.Error())
        return
    }

    // Redirect or send a success response
    c.Redirect("/some-page", 302)
}

