package controllers

import (
	"raven/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type SystemsController struct {
	beego.Controller
}

//////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////// ROUTES WITH NO HTML ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

func (c *SystemsController) Delete() {

	systemIp := c.GetString("system_ip")

	err := DeleteSystem(systemIp)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Redirect("/networks", 302) // CHANGE AS NEEDED
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////// HELPER FUNCTIONS ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

func AddSystem(system models.System) (err error) {

	o := orm.NewOrm()

	// Try and find the system in the db
	existing := models.System{Ip: system.Ip}
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

	system := models.System{Ip: systemIp}

	_, err = o.Delete(&system)
	if err != nil {
		return err
	}

	return nil
}

func GetSystem(systemIp string) (system models.System, err error) {

	o := orm.NewOrm()

	system = models.System{Ip: systemIp}
	err = o.Read(&system, "Ip")
	if err != nil {
		return models.System{}, err
	}

	return system, nil
}

func GetSystemPorts(systemIp string) (systemPorts []models.SystemPort, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.SystemPort)).RelatedSel().Filter("System__Ip", systemIp).All(&systemPorts)
	if err != nil {
		return nil, err
	}

	return systemPorts, nil
}
