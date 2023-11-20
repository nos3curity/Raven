package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type System struct {
	Ip          string        `json:"ip" orm:"pk"`
	Hostname    string        `json:"hostname"`
	Os          string        `json:"os"`
	Network     *Network      `orm:"rel(fk)" json:"network"`
	SystemPorts []*SystemPort `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(System))
}
