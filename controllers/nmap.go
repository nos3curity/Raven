package controllers

import (
	"ac-130/models"

	beego "github.com/beego/beego/v2/server/web"
)

type NmapController struct {
	beego.Controller
}

func (c *NmapController) Pwned() {
	systemIp := c.GetString("ip")
	pwnedStatus, err := c.GetBool("pwned")
	if err != nil {
		c.Ctx.WriteString("Invalid pwned status: " + err.Error())
		return
	}

	err = models.UpdateSystemPwnedStatus(systemIp, pwnedStatus)
	if err != nil {
		c.Ctx.WriteString("Error updating system: " + err.Error())
		return
	}

	c.Ctx.WriteString("System updated successfully")
}
