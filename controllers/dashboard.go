package controllers

import (
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {

	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
	c.TplName = "dashboard.html"
	return
}
