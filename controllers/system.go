package controllers

import (
	"raven/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type SystemsController struct {
	beego.Controller
}

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
