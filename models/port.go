package models

import "github.com/beego/beego/v2/client/orm"

type Port struct {
	PortNumber  int           `json:"portnumber" orm:"pk"`
	Protocol    string        `json:"protocol"`
	SystemPorts []*SystemPort `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Port))
}

func AddPort(port Port) (err error) {

	o := orm.NewOrm()

	// Try and find the system in the db
	existing := Port{PortNumber: port.PortNumber}
	readErr := o.Read(&existing, "PortNumber")

	// If no row exists, use INSERT
	if readErr == orm.ErrNoRows {

		_, err = o.Insert(&port)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return err
		}
	}

	return nil
}
