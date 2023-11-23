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

	// Get all networks
	networks, err := GetAllNetworks()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	networkSystems := make(map[string][]models.System)
	systemPorts := make(map[string][]models.SystemPort)

	// Get systems for all networks
	for _, network := range networks {
		systems, err := GetNetworkSystems(network.NetworkCidr)
		if err != nil {
			continue
		}

		networkSystems[network.NetworkCidr] = systems

		// Grab open ports for each system
		for _, system := range systems {
			ports, err := GetSystemPorts(system.Ip)
			if err != nil {
				continue
			}

			systemPorts[system.Ip] = ports
		}
	}

	c.Data["networks"] = networks
	c.Data["network_systems"] = networkSystems
	c.Data["system_ports"] = systemPorts
	c.TplName = "networks.html"
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////// ROUTES WITH NO HTML ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

func (c *NetworksController) Add() {

	networkCidr := c.GetString("network_cidr")

	teamId, err := c.GetInt("team_id")
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	err = AddNetwork(teamId, networkCidr)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Redirect("/networks", 302) // CHANGE AS NEEDED
	return
}

func (c *NetworksController) Delete() {

	networkCidr := c.GetString("network_cidr")

	err := DeleteNetwork(networkCidr)
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Redirect("/networks", 302) // CHANGE AS NEEDED
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////// HELPER FUNCTIONS ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

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

func GetNetwork(networkCidr string) (network models.Network, err error) {

	o := orm.NewOrm()
	var networks []models.Network

	_, err = o.QueryTable(new(models.Network)).RelatedSel().Filter("NetworkCidr", networkCidr).All(&networks)
	if err != nil {
		return models.Network{}, err
	}

	if len(networks) > 0 {
		network = networks[0]
	} else {
		network = models.Network{}
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

func AddNetwork(teamId int, networkCidr string) (err error) {

	o := orm.NewOrm()

	// First ensure the team exists
	team, err := GetTeam(teamId)
	if err != nil {
		return err
	}

	// Then check to see if we have a valid cidr
	_, _, err = net.ParseCIDR(networkCidr)
	if err != nil {
		return err
	}

	// Get the network ID from CIDR
	networkId, err := GetNetworkIdByCidr(networkCidr)
	if err != nil {
		return err
	}

	// Convert network ID to integer
	networkIdInt, err := ConvertIpToUint32(networkId)
	if err != nil {
		return err
	}

	// Get network broadcast from CIDR
	networkBroadcast, err := GetNetworkBroadcastByCidr(networkCidr)
	if err != nil {
		return err
	}

	// Convert broadcast to integer
	networkBroadcastInt, err := ConvertIpToUint32(networkBroadcast)
	if err != nil {
		return err
	}

	// Check if the network exists already
	exists, err := GetNetwork(networkCidr)
	if err != nil {
		return nil
	}

	// Populate the model
	network := models.Network{
		NetworkCidr:      networkCidr,
		NetworkID:        networkIdInt,
		NetworkBroadcast: networkBroadcastInt,
		Team:             &team,
	}

	if exists.NetworkCidr == "" {
		// If doesn't exist, insert a new record
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
