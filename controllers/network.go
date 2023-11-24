package controllers

import (
	"ac-130/models"

	beego "github.com/beego/beego/v2/server/web"
)

type NetworksController struct {
	beego.Controller
}

func (c *NetworksController) Add() {

	networkCidr := c.GetString("network_cidr")

	teamId, err := c.GetInt("team_id")
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	err = models.AddNetwork(teamId, networkCidr)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Redirect("/teams", 302) // CHANGE AS NEEDED
	return
}

func (c *NetworksController) Delete() {

	networkCidr := c.GetString("network_cidr")

	err := models.DeleteNetwork(networkCidr)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Redirect("/networks", 302) // CHANGE AS NEEDED
	return
}
