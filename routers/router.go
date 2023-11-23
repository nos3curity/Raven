package routers

import (
	"raven/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.DashboardController{})
	beego.Router("/nmap", &controllers.NmapController{})

	beego.Router("/teams", &controllers.TeamsController{}, "get:Setup")
	beego.Router("/teams/:id", &controllers.TeamsController{}, "get:Get")
	beego.Router("/teams/add", &controllers.TeamsController{}, "get:Add")
	beego.Router("/teams/delete", &controllers.TeamsController{}, "get:Delete")

	beego.Router("/networks/add", &controllers.NetworksController{}, "get:Add")
	beego.Router("/networks/delete", &controllers.NetworksController{}, "get:Delete")

	beego.Router("/systems/delete", &controllers.SystemsController{}, "get:Delete")

	beego.Router("/uploads", &controllers.UploadsController{}, "get:Get")
	beego.Router("/uploads/nmap", &controllers.UploadsController{}, "post:Nmap")

}
