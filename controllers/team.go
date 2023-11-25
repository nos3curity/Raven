package controllers

import (
	"fmt"
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
	systemComments := make(map[string][]models.Comment)

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

			// Grab comments for each system into a map
			comments, err := models.GetSystemComments(system.Ip)
			if err != nil {
				continue
			}
			systemComments[system.Ip] = comments
		}
	}

	// Populate the context for the template
	c.Data["teams"] = teams
	c.Data["team"] = team
	c.Data["networks"] = networks
	c.Data["network_systems"] = networkSystems
	c.Data["system_ports"] = systemPorts
	c.Data["system_comments"] = systemComments

	c.Layout = "sidebar.tpl"
	c.TplName = "team/systems.html"
	return
}

func (c *TeamsController) Add() {

	teamName := c.GetString("team_name")
	teamNetworks := c.GetStrings("network_cidr[]")

	// First add the team
	team, err := models.AddTeam(teamName)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Loop over the networks array and add them
	for _, network := range teamNetworks {
		err := models.AddNetwork(team.Id, network)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			return
		}
	}

	c.Redirect("/teams", 302) // CHANGE AS NEEDED
	return
}

func (c *TeamsController) AddMultiple() {

	var teamsNumber int
	var err error

	var teams []models.Team

	// Check to see if we have a teams_number
	teamsNumber, err = c.GetInt("teams_number")
	if (err != nil) || teamsNumber < 1 {
		c.Ctx.WriteString("No teams number provided")
		return
	}

	// Check if we have networks
	teamNetworks := c.GetStrings("network_cidr[]")
	if len(teamNetworks) < 1 {
		c.Ctx.WriteString("No networks provided")
		return
	}

	// Get the form arrays
	teamOctets := c.GetStrings("team_octet[]")
	teamIncrements := c.GetStrings("team_increment[]")

	// Check if we have the octets and increments
	if len(teamOctets) < 1 || len(teamIncrements) < 1 {
		c.Ctx.WriteString("No team octet or increments provided")
		return
	}

	// Do N loops where N is the number of teams
	for teamIncrement := 0; teamIncrement < teamsNumber; teamIncrement++ {

		fmt.Println(teamIncrement)
		// Make teams with names based on their ID
		team, err := models.AddTeam(fmt.Sprintf("Team %d", teamIncrement+1))
		if err != nil {
			c.Ctx.WriteString(err.Error())
			return
		}

		teams = append(teams, team)
	}

	// Loop over the networks array
	for i, network := range teamNetworks {

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
