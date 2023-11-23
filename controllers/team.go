package controllers

import (
	"raven/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type TeamsController struct {
	beego.Controller
}

func (c *TeamsController) Add() {

	teamName := c.GetString("team_name")

	_, err := AddTeam(teamName)
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
