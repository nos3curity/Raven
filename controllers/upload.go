package controllers

import (
	"os"
	"raven/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UploadsController struct {
	beego.Controller
}

func (c *UploadsController) Get() {
	// Get teams for the sidebar
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}

	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
	c.TplName = "upload.html"
	return
}

func (c *UploadsController) Nmap() {

	file, _, err := c.GetFile("nmap")
	if err != nil {
		c.Ctx.WriteString("Error in file upload: " + err.Error())
		return
	}
	defer file.Close()

	// Generate a random filename
	uniqueFileName := models.GenerateRandomFilename()

	// Define the full path
	tempFilePath := "uploads/" + uniqueFileName + ".xml"

	// Save the file
	err = c.SaveToFile("nmap", tempFilePath)
	if err != nil {
		c.Ctx.WriteString("Error saving file: " + err.Error())
		return
	}

	// Process the file with ParseNmap function
	err = models.ParseNmapScan(tempFilePath)
	if err != nil {
		c.Ctx.WriteString("Error parsing the nmap scan: " + err.Error())
		return
	}

	// Delete the file after processing
	err = os.Remove(tempFilePath)
	if err != nil {
		c.Ctx.WriteString("Error deleting file: " + err.Error())
		return
	}

	c.Redirect("/uploads", 302) // CHANGE AS NEEDED
}
