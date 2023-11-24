package controllers

import (
	"ac-130/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UploadsController struct {
	beego.Controller
}

func (c *UploadsController) Get() {
	// Get teams for the sidebar
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
	c.TplName = "upload.html"
	return
}
