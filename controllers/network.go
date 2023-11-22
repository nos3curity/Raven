package controllers

import (
	"fmt"
	"net"
	"raven/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/praserx/ipconv"
)

type NetworksController struct {
	beego.Controller
}

func (c *NetworksController) Get() {

	c.Ctx.WriteString("Work in Progress 🤓")
}

func (c *NetworksController) Add() {

	networkCidr := c.GetString("network_cidr")

	err := AddNetwork(networkCidr)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Ctx.WriteString("")
}

func (c *NetworksController) Delete() {

	c.Ctx.WriteString("Work in Progress 🤓")
}

/////////////////////////////////////////////////
////////////////// Helper functions go below here
/////////////////////////////////////////////////

func ConvertIpToUint32(ipString string) (intIp uint32, err error) {

	ip := net.ParseIP(ipString)
	intIp, err = ipconv.IPv4ToInt(ip)
	if err != nil {
		return 0, err
	}

	return intIp, nil
}

func GetNetworkIdByCidr(networkCidr string) (networkId string, err error) {

	_, ipNet, err := net.ParseCIDR(networkCidr)
	if err != nil {
		return "", err
	}

	networkId = ipNet.IP.String()
	return networkId, nil
}

func GetNetworkBroadcastByCidr(networkCidr string) (broadcastString string, err error) {

	_, ipNet, err := net.ParseCIDR(networkCidr)
	if err != nil {
		return "", err
	}

	networkId := ipNet.IP.To4()
	networkMask := ipNet.Mask

	broadcastInt := make(net.IP, len(networkId))
	for i := 0; i < len(networkId); i++ {
		broadcastInt[i] = networkId[i] | ^networkMask[i]
	}

	broadcastString = broadcastInt.String()

	return broadcastString, nil
}

func GetNetwork(networkCidr string) (network []models.Network, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.Network)).RelatedSel().Filter("NetworkCidr", networkCidr).All(&network)
	if err != nil {
		return nil, err
	}

	return network, nil
}

func GetAllNetworks() (networks []models.Network, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.Network)).All(&networks)
	if err != nil {
		return nil, err
	}

	return networks, nil
}

func AddNetwork(networkCidr string) (err error) {

	o := orm.NewOrm()

	_, _, err = net.ParseCIDR(networkCidr)
	if err != nil {
		return err
	}

	networkId, err := GetNetworkIdByCidr(networkCidr)
	if err != nil {
		return err
	}

	networkIdInt, err := ConvertIpToUint32(networkId)
	if err != nil {
		return err
	}

	networkBroadcast, err := GetNetworkBroadcastByCidr(networkCidr)
	if err != nil {
		return err
	}

	networkBroadcastInt, err := ConvertIpToUint32(networkBroadcast)
	if err != nil {
		return err
	}

	network := models.Network{
		NetworkCidr:      networkCidr,
		NetworkID:        networkIdInt,
		NetworkBroadcast: networkBroadcastInt,
	}

	// Check if the network exists already
	exists, err := GetNetwork(networkCidr)
	if err != nil {
		return nil
	}

	if len(exists) == 0 {
		// If not, insert a new record
		_, err = o.Insert(&network)
		if err != nil {
			return err
		}
	} else {
		// If it does, update
		_, err = o.Update(&network)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetSystemsNetwork(systemIp string) (network models.Network, err error) {

	// Convert the system ip to an integer
	intIp, err := ConvertIpToUint32(systemIp)
	if err != nil {
		return models.Network{}, err
	}

	// Get all networks
	networks, err := GetAllNetworks()
	if err != nil {
		return models.Network{}, err
	}

	// Loop over all the networks and return the one the ip belongs to
	for _, network := range networks {
		if intIp >= network.NetworkID && intIp <= network.NetworkBroadcast {
			return network, nil
		}
	}

	return models.Network{}, fmt.Errorf("Out of scope")
}

func DeleteNetwork(networkCidr string) (err error) {

	o := orm.NewOrm()

	network := models.Network{NetworkCidr: networkCidr}

	_, err = o.Delete(&network)
	if err != nil {
		return err
	}

	return nil
}

func GetNetworkSystems(networkCidr string) (systems []models.System, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(models.System)).RelatedSel().Filter("Network__NetworkCidr", networkCidr).All(&systems)
	if err != nil {
		return nil, err
	}

	return systems, nil
}
