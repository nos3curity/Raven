package routers

import (
	"raven/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/nmap", &controllers.NmapController{})

	beego.Router("/networks", &controllers.NetworksController{}, "get:Get")
	beego.Router("/networks/add", &controllers.NetworksController{}, "get:Add")
	beego.Router("/networks/delete", &controllers.NetworksController{}, "get:Delete")

	beego.Router("/systems/delete", &controllers.NetworksController{}, "get:Delete")

	beego.Router("/uploads", &controllers.UploadsController{}, "get:Get")
	beego.Router("/uploads/nmap", &controllers.UploadsController{}, "post:Nmap")
}
