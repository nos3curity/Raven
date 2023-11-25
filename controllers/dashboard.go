package controllers

import (
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {

	// Get all teams
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	// Pass to the template
	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
	c.TplName = "dashboard.html"
	return
}
