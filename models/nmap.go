package models

import (
	"fmt"
	"os"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/lair-framework/go-nmap"
)

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

		// Add the scan timestamp to the system
		system.LatestScan = (*time.Time)(&scan.RunStats.Finished.Time)

		// Check if we have system already
		existingSystem, err := GetSystem(system.Ip)
		if (err != nil) && (err != orm.ErrNoRows) {
			fmt.Println(err)
			continue
		}

		// Check if we have the latest scan and skip the host if we do
		if system.LatestScan != nil && existingSystem.LatestScan != nil {
			if system.LatestScan.Before(*existingSystem.LatestScan) || system.LatestScan.Equal(*existingSystem.LatestScan) {
				continue
			}
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

func ParseNmapHost(host nmap.Host) (system System, err error) {

	if len(host.Addresses) == 0 {
		return System{}, fmt.Errorf("Host has no IP addresses")
	}

	if host.Addresses[0].AddrType != "ipv4" {
		return System{}, fmt.Errorf("Not an IPv4 address")
	}

	var ipAddress string
	var hostname string
	var fingerprint string
	var osFamily string

	// Assign host values to variables
	ipAddress = host.Addresses[0].Addr

	if len(host.Hostnames) != 0 {
		hostname = host.Hostnames[0].Name
	} else {
		hostname = ""
	}

	if len(host.Os.OsMatches) > 0 {
		fingerprint = host.Os.OsMatches[0].Name
		osFamily = host.Os.OsMatches[0].OsClasses[0].OsFamily
	} else {
		fingerprint = ""
		osFamily = ""
	}

	// Find which network the system belongs to
	network, err := GetSystemsNetwork(ipAddress)
	if err != nil {
		return System{}, err
	}

	// Assign the variables to the model
	system = System{
		Ip:       ipAddress,
		Hostname: hostname,
		OsFamily: osFamily,
		Os:       fingerprint,
		Network:  &network,
	}

	return system, nil
}

func ParseNmapPort(port nmap.Port) (openPort Port, err error) {

	if port.State.State != "open" {
		return Port{}, fmt.Errorf("Port closed or filtered")
	}

	var portNumber int
	var protocol string
	var serviceName string
	var serviceVersion string
	var serviceProduct string

	// Assign defaults
	serviceProduct = "N/A"
	serviceName = "N/A"
	serviceVersion = "N/A"

	// Only assign values if not empty
	if port.Service.Name != "" {
		serviceName = port.Service.Name
	}
	if port.Service.Version != "" {
		serviceVersion = port.Service.Version
	}
	if port.Service.Product != "" {
		serviceProduct = port.Service.Product
	}

	// Assign values to variables
	portNumber = port.PortId
	protocol = port.Protocol

	// Assign the variables to the model
	openPort = Port{
		PortNumber:         portNumber,
		Protocol:           protocol,
		PortServiceName:    serviceName,
		PortServiceVersion: serviceVersion,
		PortServiceProduct: serviceProduct,
	}

	return openPort, nil
}
