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
	SystemPorts []*SystemPort `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(System))
}
