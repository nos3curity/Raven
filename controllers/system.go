package controllers

import (
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type SystemsController struct {
	beego.Controller
}

func (c *SystemsController) Add() {

	systemIp := c.GetString("system_ip")
	hostname := c.GetString("hostname")
	osFamily := c.GetString("os_family")   // Linux, Windows, Other
	osVersion := c.GetString("os_version") // Windows Server 2019
	network_cidr := c.GetString("network_cidr")

	if (systemIp == "") || (network_cidr == "") {
		c.Ctx.WriteString("Provide at least the IP and network_cidr ")
		return
	}

	network, err := models.GetNetwork(network_cidr)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	system := models.System{
		Ip:       systemIp,
		Hostname: hostname,
		OsFamily: osFamily,
		Os:       osVersion,
		Network:  &network,
	}

	err = models.AddSystem(system)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	return
}

func (c *SystemsController) SetOs() {

	systemIp := c.GetString("system_ip")
	osFamily := c.GetString("os_family")   // Linux, Windows, Other
	osVersion := c.GetString("os_version") // Windows Server 2019

	// Handle empty input
	if (systemIp == "") || (osFamily == "") || (osVersion == "") {

		c.Ctx.WriteString("Provide IP, os family and os version")
		return
	}

	_, err := models.GetSystem(systemIp)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	err = models.SetSystemOs(systemIp, osFamily, osVersion)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	return
}

func (c *SystemsController) SetHostname() {

	systemIp := c.GetString("system_ip")
	hostname := c.GetString("hostname")

	// Handle empty input
	if (systemIp == "") || (hostname == "") {

		c.Ctx.WriteString("Provide an IP and hostname")
		return
	}

	_, err := models.GetSystem(systemIp)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	err = models.SetSystemHostname(systemIp, hostname)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	return
}

func (c *SystemsController) Delete() {

	systemIp := c.GetString("system_ip")
	err := models.DeleteSystem(systemIp)

	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	c.Redirect("/", 302) // CHANGE AS NEEDED
	return
}
