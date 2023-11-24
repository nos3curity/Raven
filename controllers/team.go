package controllers

import (
	"raven/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type TeamsController struct {
	beego.Controller
}

func (c *TeamsController) Setup() {

	teamNetworks := make(map[int][]models.Network)

	// Get all teams
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	// Fetch the team networks
	for _, team := range teams {

		networks, err := models.GetTeamNetworks(team.Id)
		if err != nil {
			c.Ctx.WriteString(err.Error())
		}

		teamNetworks[team.Id] = networks

	}

	c.Data["team_networks"] = teamNetworks
	c.Data["teams"] = teams

	c.Layout = "sidebar.tpl"
	c.TplName = "team/setup.html"
}

func (c *TeamsController) Get() {

	// Parse the team ID integer
	teamId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	networkSystems := make(map[string][]models.System)
	systemPorts := make(map[string][]models.SystemPort)

	// Get teams for the sidebar
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// First, fetch the team information
	team, err := models.GetTeam(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Fetch the team networks
	networks, err := models.GetTeamNetworks(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get systems for all networks
	for _, network := range networks {
		systems, err := models.GetNetworkSystems(network.NetworkCidr)
		if err != nil {
			continue
		}

		networkSystems[network.NetworkCidr] = systems

		// Grab open ports for each system
		for _, system := range systems {
			ports, err := models.GetSystemPorts(system.Ip)
			if err != nil {
				continue
			}

			systemPorts[system.Ip] = ports
		}
	}

	// Populate the context for the template
	c.Data["teams"] = teams
	c.Data["team"] = team
	c.Data["networks"] = networks
	c.Data["network_systems"] = networkSystems
	c.Data["system_ports"] = systemPorts

	c.Layout = "sidebar.tpl"
	c.TplName = "team/systems.html"
	return
}

func (c *TeamsController) Add() {

	teamName := c.GetString("team_name")

	_, err := models.AddTeam(teamName)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	c.Redirect("/teams", 302) // CHANGE AS NEEDED
	return
}

func (c *TeamsController) Delete() {

	teamId, err := c.GetInt("team_id")
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	err = models.DeleteTeam(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Redirect("/teams", 302) // CHANGE AS NEEDED
	return
}
