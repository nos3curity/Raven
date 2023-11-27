package controllers

import (
	"raven/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type NetworksController struct {
	beego.Controller
}

func (c *NetworksController) Prepare() {
	sidebar := &SiderbarController{Controller: c.Controller}
	sidebar.GetTeams()
}

func (c *NetworksController) Add() {

	// Parse the team id
	teamId, err := c.GetInt("team_id")
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Try to find the team
	_, err = models.GetTeam(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get the network cidr
	networkCidr := c.GetString("network_cidr")
	if networkCidr == "" {
		c.Ctx.WriteString("No network cidr provided")
		return
	}

	// Add the network
	err = models.AddNetwork(teamId, networkCidr)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	c.Redirect("/teams", 302) // CHANGE AS NEEDED
	return
}

func (c *NetworksController) AddMultiple() {

	networks := c.GetStrings("network_cidr[]")

	if len(networks) < 1 {
		c.Ctx.WriteString("No networks provided")
		return
	}

	// Get all teams
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	// Get the form arrays
	teamOctets := c.GetStrings("team_octet[]")
	teamIncrements := c.GetStrings("team_increment[]")

	if len(teamOctets) < 1 || len(teamIncrements) < 1 {
		c.Ctx.WriteString("No team octet or increments provided")
		return
	}

	// Loop over the networks array
	for i, network := range networks {

		teamOctet, err := strconv.Atoi(teamOctets[i])
		teamIncrement, err := strconv.Atoi(teamIncrements[i])
		if err != nil {
			c.Ctx.WriteString(err.Error())
			return
		}

		for i := 0; i < len(teams); i++ {

			// Add the network
			err = models.AddNetwork(teams[i].Id, network)
			if err != nil {
				c.Ctx.WriteString(err.Error())
				return
			}

			// Increment the network
			network, err = models.IncrementNetwork(network, teamOctet, teamIncrement)
			if err != nil {
				c.Ctx.WriteString(err.Error())
				return
			}
		}
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
