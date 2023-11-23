package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Get() {

	teams, err := GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
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
