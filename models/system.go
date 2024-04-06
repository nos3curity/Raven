package models

import (
	"sort"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type System struct {
	Ip          string        `json:"ip" orm:"pk"`
	Hostname    string        `json:"hostname"`
	OsFamily    string        `json:"os_family"`
	Os          string        `json:"os"`
	Network     *Network      `orm:"rel(fk);column(network);on_delete(cascade)"`
	SystemPorts []*SystemPort `orm:"reverse(many);on_delete(cascade)"`
	LatestScan  *time.Time    `orm:"null" json:"latest_scan"`
	Pwned       bool          `json:"pwned"`
}

func init() {
	orm.RegisterModel(new(System))
}

func AddSystem(system System) (err error) {

	o := orm.NewOrm()

	// Try and find the system in the db
	existing := System{Ip: system.Ip}
	readErr := o.Read(&existing, "Ip")

	// If no row exists, use INSERT
	if readErr == orm.ErrNoRows {

		_, err = o.Insert(&system)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return err
		}
		// If a row exists, use UPDATE
	} else if readErr == nil {

		_, err = o.Update(&system)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return err
		}
	}

	return nil
}

func DeleteSystem(systemIp string) (err error) {

	o := orm.NewOrm()

	system := System{Ip: systemIp}

	_, err = o.Delete(&system)
	if err != nil {
		return err
	}

	return nil
}

func GetSystem(systemIp string) (system System, err error) {

	o := orm.NewOrm()

	system = System{Ip: systemIp}
	err = o.Read(&system, "Ip")
	if err != nil {
		return System{}, err
	}

	return system, nil
}

func GetSystemPorts(systemIp string) (systemPorts []SystemPort, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(SystemPort)).
		RelatedSel().
		Filter("System__Ip", systemIp).
		All(&systemPorts)
	if err != nil {
		return nil, err
	}

	// Sort the systemPorts array by port number
	sort.Slice(systemPorts, func(i, j int) bool {
		return systemPorts[i].Port.PortNumber < systemPorts[j].Port.PortNumber
	})

	return systemPorts, nil
}

func UpdateSystemPwnedStatus(systemIp string, pwned bool) (err error) {
	o := orm.NewOrm()

	system := System{Ip: systemIp}
	err = o.Read(&system, "Ip")
	if err != nil {
		return err
	}

	system.Pwned = pwned

	_, err = o.Update(&system)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

func SetSystemHostname(systemIp string, hostname string) (err error) {

	o := orm.NewOrm()

	system, err := GetSystem(systemIp)
	if err != nil {
		return err
	}

	system.Hostname = hostname
	_, err = o.Update(&system)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

func SetSystemOs(systemIp string, osFamily string, osVersion string) (err error) {

	o := orm.NewOrm()

	system, err := GetSystem(systemIp)
	if err != nil {
		return err
	}

	system.OsFamily = osFamily
	system.Os = osVersion

	_, err = o.Update(&system)
	if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
		return err
	}

	return nil
}

func GetSystemsTeam(systemIp string) (team *Team, err error) {

	// Get the system
	system, err := GetSystem(systemIp)
	if err != nil {
		return nil, err
	}

	// Get the system network
	networkCidr := system.Network.NetworkCidr
	network, err := GetNetwork(networkCidr)
	if err != nil {
		return nil, err
	}

	team = network.Team

	return team, nil
}
