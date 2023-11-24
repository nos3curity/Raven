package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type System struct {
	Ip          string        `json:"ip" orm:"pk"`
	Hostname    string        `json:"hostname"`
	OsFamily    string        `json:"os_family"`
	Os          string        `json:"os"`
	Network     *Network      `orm:"rel(fk);column(network);on_delete(cascade)"`
	SystemPorts []*SystemPort `orm:"reverse(many);on_delete(cascade)"`
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
		if err != nil {
			return err
		}
		// If a row exists, use UPDATE
	} else if readErr == nil {

		_, err = o.Update(&system)
		if err != nil {
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

	_, err = o.QueryTable(new(SystemPort)).RelatedSel().Filter("System__Ip", systemIp).All(&systemPorts)
	if err != nil {
		return nil, err
	}

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
	return err
}

func SetSystemHostname(systemIp string, hostname string) (err error) {

	o := orm.NewOrm()

	system, err := GetSystem(systemIp)
	if err != nil {
		return err
	}

	system.Hostname = hostname
	o.Update(&system)

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
	o.Update(&system)

	return nil
}
