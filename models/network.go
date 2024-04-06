package models

import (
	"bytes"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/praserx/ipconv"
)

type Network struct {
	NetworkCidr      string    `json:"network_cidr" orm:"pk"`
	NetworkID        uint32    `json:"network_id"`
	NetworkBroadcast uint32    `json:"network_broadcast"`
	NetworkSystems   []*System `orm:"reverse(many)"`
	Team             *Team     `orm:"rel(fk);on_delete(cascade)" json:"team"`
}

func init() {
	orm.RegisterModel(new(Network))
}

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

func GetNetwork(networkCidr string) (network Network, err error) {

	o := orm.NewOrm()
	var networks []Network

	_, err = o.QueryTable(new(Network)).RelatedSel().Filter("NetworkCidr", networkCidr).All(&networks)
	if err != nil {
		return Network{}, err
	}

	if len(networks) > 0 {
		network = networks[0]
	} else {
		network = Network{}
	}

	return network, nil
}

func GetAllNetworks() (networks []Network, err error) {

	o := orm.NewOrm()

	_, err = o.QueryTable(new(Network)).All(&networks)
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
	network := Network{
		NetworkCidr:      networkCidr,
		NetworkID:        networkIdInt,
		NetworkBroadcast: networkBroadcastInt,
		Team:             &team,
	}

	if exists.NetworkCidr == "" {
		// If doesn't exist, insert a new record
		_, err = o.Insert(&network)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return err
		}
	} else {
		// If it does, update
		_, err = o.Update(&network)
		if (err != nil) && (err != orm.ErrLastInsertIdUnavailable) {
			return err
		}
	}

	return nil
}

func GetSystemsNetwork(systemIp string) (network Network, err error) {

	// Convert the system ip to an integer
	intIp, err := ConvertIpToUint32(systemIp)
	if err != nil {
		return Network{}, err
	}

	// Get all networks
	networks, err := GetAllNetworks()
	if err != nil {
		return Network{}, err
	}

	// Loop over all the networks and return the one the ip belongs to
	for _, network := range networks {
		if intIp >= network.NetworkID && intIp <= network.NetworkBroadcast {
			return network, nil
		}
	}

	return Network{}, fmt.Errorf("Out of scope")
}

func DeleteNetwork(networkCidr string) (err error) {

	o := orm.NewOrm()

	network := Network{NetworkCidr: networkCidr}

	_, err = o.Delete(&network)
	if err != nil {
		return err
	}

	return nil
}

func GetNetworkSystems(networkCidr string) (systems []System, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(System)).
		RelatedSel().
		Filter("Network__NetworkCidr", networkCidr).
		All(&systems)
	if err != nil {
		return nil, err
	}

	// Sort the systems array by IP address
	sort.Slice(systems, func(i, j int) bool {
		ip1 := net.ParseIP(systems[i].Ip).To4()
		ip2 := net.ParseIP(systems[j].Ip).To4()
		return bytes.Compare(ip1, ip2) < 0
	})

	return systems, nil
}

func ConvertNetmaskToCidr(mask net.IPMask) int {
	size, _ := mask.Size()
	return size
}

func IncrementNetwork(networkCidr string, octet int, increment int) (newNetworkCidr string, err error) {
	// Parse the CIDR
	ip, ipnet, err := net.ParseCIDR(networkCidr)
	if err != nil {
		return "", err
	}

	// Split the IP into octets
	ipOctets := strings.Split(ip.String(), ".")

	// Increment the specified octet
	if octet < 1 || octet > 4 {
		return "", fmt.Errorf("invalid octet: %d", octet)
	}
	octetIndex := octet - 1 // Adjust for 0-based indexing
	octetValue, err := strconv.Atoi(ipOctets[octetIndex])
	if err != nil {
		return "", err
	}

	octetValue += increment
	if octetValue > 255 {
		return "", fmt.Errorf("octet value out of range after increment")
	}
	ipOctets[octetIndex] = strconv.Itoa(octetValue)

	// Reassemble the IP
	newIP := net.ParseIP(strings.Join(ipOctets, "."))
	if newIP == nil {
		return "", fmt.Errorf("failed to parse new IP")
	}

	// Construct new CIDR
	newNetworkCidr = fmt.Sprintf("%s/%d", newIP, ConvertNetmaskToCidr(ipnet.Mask))
	return newNetworkCidr, nil
}

func GetNetworkNmapTimestamp(networkCidr string) (*time.Time, error) {
	systems, err := GetNetworkSystems(networkCidr)
	if err != nil {
		return nil, err
	}

	if len(systems) == 0 {
		return nil, fmt.Errorf("no systems found for the network")
	}

	// Sort systems by LatestScan in descending order to find the latest scan
	sort.Slice(systems, func(i, j int) bool {
		if systems[i].LatestScan == nil {
			return false
		}
		if systems[j].LatestScan == nil {
			return true
		}
		return systems[i].LatestScan.After(*systems[j].LatestScan)
	})

	// The first system should now have the latest LastScanned value
	for _, system := range systems {
		if system.LatestScan != nil {
			return system.LatestScan, nil
		}
	}

	// If all systems have a nil LatestScan, return an error
	return nil, fmt.Errorf("no latest scan timestamp found")
}
