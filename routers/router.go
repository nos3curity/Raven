package routers

import (
	"ac-130/controllers"

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

	// Teams functionality
	beego.Router("/teams", &controllers.TeamsController{}, "get:Setup")
	beego.Router("/teams/:id", &controllers.TeamsController{}, "get:Get")
	beego.Router("/teams/add", &controllers.TeamsController{}, "get:Add")
	beego.Router("/teams/delete", &controllers.TeamsController{}, "get:Delete")

	// Networks functionality
	beego.Router("/networks/add", &controllers.NetworksController{}, "get:Add")
	beego.Router("/networks/delete", &controllers.NetworksController{}, "get:Delete")

	// Networks functionality
	beego.Router("/systems/delete", &controllers.SystemsController{}, "get:Delete")

	// Uploads functionality
	beego.Router("/uploads", &controllers.UploadsController{}, "get:Get")

	// External API functionality
	beego.Router("/api/nmap", &controllers.ApiController{}, "post:Nmap")
	beego.Router("/api/pwned", &controllers.ApiController{}, "post:Pwned")

}
