package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func GetErrorMessage(err error) string {

	defaultError := "🤓🤓🤓 OOPSIE WOOPSIE!! Uwu We make a fucky wucky!! A wittle fucko boingo! The code monkeys at our headquarters are working VEWY HAWD to fix this! 🤓🤓🤓"

	runmode, err := beego.AppConfig.String("runmove")
	if err != nil {
		fmt.Println("something is really fucking wrong")
	}

	if runmode == "dev" {
		return err.Error()
	} else {
		return defaultError
	}

}
