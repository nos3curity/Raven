package controllers

import (
	"ac-130/models"

	beego "github.com/beego/beego/v2/server/web"
)

type SystemsController struct {
	beego.Controller
}

func (c *SystemsController) Delete() {

	systemIp := c.GetString("system_ip")
	err := models.DeleteSystem(systemIp)

	if err != nil {
		c.Ctx.WriteString("fucked up")
		return
	}

	c.Redirect("/", 302) // CHANGE AS NEEDED
	return
}
