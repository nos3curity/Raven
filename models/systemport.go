package models

import "github.com/beego/beego/v2/client/orm"

type SystemPort struct {
	Id     int     `json:"id"`
	System *System `orm:"rel(fk);column(system_ip)"`
	Port   *Port   `orm:"rel(fk);column(port_number)"`
	Open   bool    `json:"open"`
}

func init() {
	orm.RegisterModel(new(SystemPort))
}
