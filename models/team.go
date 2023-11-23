package models

import "github.com/beego/beego/v2/client/orm"

type Team struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Networks []*Network `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Team))
}
