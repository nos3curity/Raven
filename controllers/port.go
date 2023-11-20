package controllers

import (
	"raven/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type PortsController struct {
	beego.Controller
}

func AddPort(port models.Port) (err error) {

	o := orm.NewOrm()

	// Try and find the system in the db
	existing := models.Port{PortNumber: port.PortNumber}
	readErr := o.Read(&existing, "PortNumber")

	// If no row exists, use INSERT
	if readErr == orm.ErrNoRows {

		_, err = o.Insert(&port)
		if err != nil {
			return err
		}
	}

	return nil
}
