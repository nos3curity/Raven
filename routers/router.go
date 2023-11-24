package routers

import (
	"raven/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.InsertFilter("/*", beego.BeforeRouter, controllers.ValidateJwtFilter)
	beego.Router("/", &controllers.DashboardController{})

	beego.Router("/login", &controllers.LoginController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "post:SignIn")

	beego.Router("/teams", &controllers.TeamsController{}, "get:Setup")
	beego.Router("/teams/:id", &controllers.TeamsController{}, "get:Get")
	beego.Router("/teams/add", &controllers.TeamsController{}, "get:Add")
	beego.Router("/teams/delete", &controllers.TeamsController{}, "get:Delete")

	beego.Router("/networks/add", &controllers.NetworksController{}, "get:Add")
	beego.Router("/networks/delete", &controllers.NetworksController{}, "get:Delete")

	beego.Router("/systems/delete", &controllers.SystemsController{}, "get:Delete")

	beego.Router("/uploads", &controllers.UploadsController{}, "get:Get")
	beego.Router("/uploads/nmap", &controllers.UploadsController{}, "post:Nmap")
	beego.Router("/pwned", &controllers.NmapController{}, "post:Pwned")

}
