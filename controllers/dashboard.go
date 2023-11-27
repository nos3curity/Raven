package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DashboardController struct {
	beego.Controller
}

func (c *DashboardController) Prepare() {
	sidebar := &SiderbarController{Controller: c.Controller}
	sidebar.GetTeams()
}

func (c *DashboardController) Get() {

	c.Layout = "layout/sidebar.tpl"
	c.TplName = "dashboard.html"
	return
}
