package controllers

import (
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UploadsController struct {
	beego.Controller
}

func (c *UploadsController) Prepare() {
	sidebar := &SiderbarController{Controller: c.Controller}
	sidebar.GetTeams()
}

func (c *UploadsController) Get() {

	c.Data["loot_tags"] = models.LootTags
	c.Layout = "layout/sidebar.tpl"
	c.TplName = "upload.html"
	return
}
