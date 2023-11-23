package controllers

import (
	"raven/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type TeamsController struct {
	beego.Controller
}

func (c *TeamsController) Setup() {

	teamNetworks := make(map[int][]models.Network)

	// Get all teams
	teams, err := GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	// Fetch the team networks
	for _, team := range teams {

		networks, err := GetTeamNetworks(team.Id)
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
	teams, err := GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// First, fetch the team information
	team, err := GetTeam(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Fetch the team networks
	networks, err := GetTeamNetworks(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get systems for all networks
	for _, network := range networks {
		systems, err := GetNetworkSystems(network.NetworkCidr)
		if err != nil {
			continue
		}

		networkSystems[network.NetworkCidr] = systems

		// Grab open ports for each system
		for _, system := range systems {
			ports, err := GetSystemPorts(system.Ip)
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

	_, err := AddTeam(teamName)
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

	err = DeleteTeam(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Redirect("/teams", 302) // CHANGE AS NEEDED
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////// HELPER FUNCTIONS ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

func AddTeam(teamName string) (team models.Team, err error) {

	o := orm.NewOrm()

	team = models.Team{
		Name: teamName,
	}

	_, err = o.Insert(&team)
	if err != nil {
		return models.Team{}, err
	}

	return team, nil
}

func GetTeam(teamId int) (team models.Team, err error) {

	o := orm.NewOrm()
	var teams []models.Team

	_, err = o.QueryTable(new(models.Team)).RelatedSel().Filter("Id", teamId).All(&teams)
	if err != nil {
		return models.Team{}, err
	}

	team = teams[0]

	return team, nil
}

func GetAllTeams() (teams []models.Team, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.Team)).All(&teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func DeleteTeam(teamId int) (err error) {

	o := orm.NewOrm()

	team := models.Team{Id: teamId}

	_, err = o.Delete(&team)
	if err != nil {
		return err
	}

	return nil
}

func RenameTeam(teamId int, teamName string) (err error) {

	o := orm.NewOrm()

	team, err := GetTeam(teamId)
	if err != nil {
		return err
	}

	team.Name = teamName
	o.Update(&team)

	return nil
}

func GetTeamNetworks(teamId int) (networks []models.Network, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.Network)).RelatedSel().Filter("Team__Id", teamId).All(&networks)
	if err != nil {
		return nil, err
	}

	return networks, nil
}
