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
	referrer := c.Ctx.Request.Referer()
	if referrer == "" {
		referrer = "/teams" // Replace with your default page
	}
	c.Redirect(referrer, 302)
}

// Delete handles the deletion of a comment
func (c *CommentsController) Delete() {
	// Retrieve the comment ID from the query parameter
	commentId, err := c.GetInt("comment_id")
	if err != nil {
		c.Ctx.WriteString("Invalid comment ID")
		return
	}

	// Call the DeleteComment function from the models package
	err = models.DeleteComment(commentId)
	if err != nil {
		c.Ctx.WriteString("Error deleting comment: " + err.Error())
		return
	}

	// Redirect to a relevant page after deletion
	referrer := c.Ctx.Request.Referer()
	if referrer == "" {
		referrer = "/teams" // Replace with your default page
	}
	c.Redirect(referrer, 302)
}
