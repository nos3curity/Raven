package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Network struct {
	NetworkCidr      string    `json:"network_cidr" orm:"pk"`
	NetworkID        uint32    `json:"network_id"`
	NetworkBroadcast uint32    `json:"network_broadcast"`
	NetworkSystems   []*System `orm:"reverse(many)"`
	//Team             *Team  `orm:"rel(fk);on_delete(cascade)" json:"team"`
}

func init() {
	orm.RegisterModel(new(Network))
}
