package controllers

import (
	"fmt"
	"os"
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/lair-framework/go-nmap"
)

type NmapController struct {
	beego.Controller
}

func (c *NmapController) Get() {
	err := ParseNmapScan("big.xml")
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString("Work in Progress 🤓")
}

func ParseNmapScan(scanPath string) (err error) {

	// Read the XML file
	data, err := os.ReadFile(scanPath)
	if err != nil {
		return err
	}

	// Parse the Nmap scan results
	var scan *nmap.NmapRun
	scan, err = nmap.Parse(data)
	if err != nil {
		return err
	}

	// Iterate over scan results
	for _, host := range scan.Hosts {

		// Parse the XML host to the system model
		system, err := ParseNmapHost(host)
		if err != nil {
			continue

		}

		// Add the system to the database
		err = AddSystem(system)
		if err != nil {
			continue
		}

		err = SetAllSystemPortsClosedByIp(system.Ip)
		if err != nil {
			continue
		}

		for _, port := range host.Ports {

			// Parse the port from XML into a model
			openPort, err := ParseNmapPort(port)
			if err != nil {
				continue
			}

			// Add the port to the ports database
			err = AddPort(openPort)
			if err != nil {
				continue
			}

			// Add the systemport association
			err = AddOpenSystemPort(&system, &openPort)
			if err != nil {
				continue
			}
		}
	}

	return nil
}

func ParseNmapHost(host nmap.Host) (system models.System, err error) {

	if len(host.Addresses) == 0 {
		return models.System{}, fmt.Errorf("Host has no IP addresses")
	}

	if host.Addresses[0].AddrType != "ipv4" {
		return models.System{}, fmt.Errorf("Not an IPv4 address")
	}

	var ipAddress string
	var hostname string
	var fingerprint string

	// Assign host values to variables
	ipAddress = host.Addresses[0].Addr

	if len(host.Hostnames) != 0 {
		hostname = host.Hostnames[0].Name
	} else {
		hostname = ""
	}

	if len(host.Os.OsFingerprints) != 0 {
		fingerprint = host.Os.OsFingerprints[0].Fingerprint
	} else {
		fingerprint = ""
	}

	// Find which network the system belongs to
	network, err := GetSystemsNetwork(ipAddress)
	if err != nil {
		return models.System{}, err
	}

	// Assign the variables to the model
	system = models.System{
		Ip:       ipAddress,
		Hostname: hostname,
		Os:       fingerprint,
		Network:  &network,
	}

	return system, nil
}

func ParseNmapPort(port nmap.Port) (openPort models.Port, err error) {

	if port.State.State != "open" {
		return models.Port{}, fmt.Errorf("Port closed or filtered")
	}

	var portNumber int
	var protocol string

	// Assign values to variables
	portNumber = port.PortId
	protocol = port.Protocol

	// Assign the variables to the model
	openPort = models.Port{
		PortNumber: portNumber,
		Protocol:   protocol,
	}

	return openPort, nil
}
