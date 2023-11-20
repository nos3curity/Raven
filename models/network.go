package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Network struct {
	NetworkCidr      string `json:"network_cidr" orm:"pk"`
	NetworkID        uint32 `json:"network_id"`
	NetworkBroadcast uint32 `json:"network_broadcast"`
	// Team *Team
}

func init() {
	orm.RegisterModel(new(Network))
}
