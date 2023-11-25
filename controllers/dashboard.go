package controllers

import (
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {

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

	// Get all networks
	networks, err := models.GetAllNetworks()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	networkSystems := make(map[string][]models.System)
	systemPorts := make(map[string][]models.SystemPort)

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

	c.Data["networks"] = networks
	c.Data["network_systems"] = networkSystems
	c.Data["system_ports"] = systemPorts

	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
	c.TplName = "dashboard.html"
	return
}
