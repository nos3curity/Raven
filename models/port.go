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
