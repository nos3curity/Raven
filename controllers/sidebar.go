package controllers

import (
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type SiderbarController struct {
	beego.Controller
}

func (c *SiderbarController) GetTeams() {
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	c.Data["teams"] = teams
}
