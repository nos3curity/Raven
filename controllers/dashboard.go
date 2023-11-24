package controllers

import (
	"fmt"
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
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
	c.TplName = "dashboard.html"
	return
}

func GetErrorMessage(err error) string {

	defaultError := "🤓🤓🤓 OOPSIE WOOPSIE!! Uwu We make a fucky wucky!! A wittle fucko boingo! The code monkeys at our headquarters are working VEWY HAWD to fix this! 🤓🤓🤓"

	runmode, err := beego.AppConfig.String("runmode")
	if err != nil {
		fmt.Println("something is really fucking wrong")
	}

	if runmode == "dev" {
		return err.Error()
	} else {
		return defaultError
	}

}
