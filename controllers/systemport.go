package controllers

import (
	"raven/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type SystemPortsController struct {
	beego.Controller
}

func AddOpenSystemPort(system *models.System, port *models.Port) (err error) {

	o := orm.NewOrm()

	systemPort := models.SystemPort{
		System: system,
		Port:   port,
		Open:   true,
	}

	// Check if there are any port associations for this system
	systemPorts, err := GetSystemPortsByIp(system.Ip)
	if len(systemPorts) == 0 {

		// If no associations exist, create it
		_, err = o.Insert(&systemPort)
		if err != nil {
			return err
		}
	} else {

		// See if this particular association exists
		systemPorts, err = GetSystemPort(system.Ip, port.PortNumber)
		if err != nil {
			return err
		}

		if len(systemPorts) == 0 {

			// If no associations exist, create it
			_, err = o.Insert(&systemPort)
			if err != nil {
				return err
			}
		} else {

			// If it exists, update it
			systemPort.Id = systemPorts[0].Id
			_, err = o.Update(&systemPort)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetSystemPort(systemIp string, port int) (systemPorts []models.SystemPort, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.SystemPort)).RelatedSel().Filter("System__Ip", systemIp).Filter("Port__PortNumber", port).All(&systemPorts)
	if err != nil {
		return nil, err
	}

	return systemPorts, nil
}

func GetSystemPortsByIp(systemIp string) (systemPorts []models.SystemPort, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.SystemPort)).RelatedSel().Filter("System__Ip", systemIp).All(&systemPorts)
	if err != nil {
		return nil, err
	}

	return systemPorts, nil
}

func SetAllSystemPortsClosedByIp(systemIp string) (err error) {

	o := orm.NewOrm()

	// Update Open field to false for all records matching the system IP
	_, err = o.QueryTable(new(models.SystemPort)).Filter("System__Ip", systemIp).Update(orm.Params{
		"Open": false,
	})

	if err != nil {
		return err
	}

	return nil
}
