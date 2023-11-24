package routers

import (
	"raven/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// Authentication filter
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.ValidateJwtFilter)

	// Dashboard
	beego.Router("/", &controllers.DashboardController{})

	// Login functionality
	beego.Router("/login", &controllers.LoginController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "post:SignIn")
	beego.Router("/profile", &controllers.LoginController{}, "get:Profile")

	// Teams functionality
	beego.Router("/teams", &controllers.TeamsController{}, "get:Setup")
	beego.Router("/teams/:id", &controllers.TeamsController{}, "get:Get")
	beego.Router("/teams/add", &controllers.TeamsController{}, "get:Add")       // TODO: convert to POST
	beego.Router("/teams/delete", &controllers.TeamsController{}, "get:Delete") // TODO: convert to POST

	// Networks functionality
	beego.Router("/networks/add", &controllers.NetworksController{}, "get:Add")       // TODO: convert to POST
	beego.Router("/networks/delete", &controllers.NetworksController{}, "get:Delete") // TODO: convert to POST

	// Systems functionality
	beego.Router("/systems/add", &controllers.SystemsController{}, "post:Add")
	beego.Router("/systems/os", &controllers.SystemsController{}, "post:SetOs")
	beego.Router("/systems/hostname", &controllers.SystemsController{}, "post:SetHostname")
	beego.Router("/systems/delete", &controllers.SystemsController{}, "post:Delete")

	// Uploads functionality
	beego.Router("/uploads", &controllers.UploadsController{}, "get:Get")

	// External API functionality
	beego.Router("/api/nmap", &controllers.ApiController{}, "post:Nmap")
	beego.Router("/api/pwned", &controllers.ApiController{}, "post:Pwned")

}
