package controllers

import (
	"ac-130/models"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Pwned() {
	systemIp := c.GetString("ip")
	pwnedStatus, err := c.GetBool("pwned")
	if err != nil {
		c.Ctx.WriteString("Invalid pwned status: " + err.Error())
		return
	}

	err = models.UpdateSystemPwnedStatus(systemIp, pwnedStatus)
	if err != nil {
		c.Ctx.WriteString("Error updating system: " + err.Error())
		return
	}

	c.Ctx.WriteString("System updated successfully")
}

func (c *ApiController) Nmap() {

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
