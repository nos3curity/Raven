package routers

import (
	"raven/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/nmap", &controllers.NmapController{})
	beego.Router("/networks/add", &controllers.NetworksController{})
}
