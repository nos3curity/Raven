package models

import "github.com/beego/beego/v2/client/orm"

type Team struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Networks []*Network `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Team))
}

func AddTeam(teamName string) (team Team, err error) {

	o := orm.NewOrm()

	team = Team{
		Name: teamName,
	}

	_, err = o.Insert(&team)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return Team{}, err
	}

	return team, nil
}

func GetTeam(teamId int) (team Team, err error) {

	o := orm.NewOrm()
	var teams []Team

	_, err = o.QueryTable(new(Team)).RelatedSel().Filter("Id", teamId).All(&teams)
	if err != nil {
		return Team{}, err
	}

	team = teams[0]

	return team, nil
}

func GetAllTeams() (teams []Team, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(Team)).All(&teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func DeleteTeam(teamId int) (err error) {

	o := orm.NewOrm()

	team := Team{Id: teamId}

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

	_, err = o.Update(&team)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

func GetTeamNetworks(teamId int) (networks []Network, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(Network)).RelatedSel().Filter("Team__Id", teamId).All(&networks)
	if err != nil {
		return nil, err
	}

	return networks, nil
}
