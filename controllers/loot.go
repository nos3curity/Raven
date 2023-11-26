package controllers

import (
	"os"
	"raven/models"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	beego "github.com/beego/beego/v2/server/web"
)

type LootController struct {
	beego.Controller
}

func (c *LootController) Get() {

	teamLootedSystems := make(map[int][]string)

	// Get all teams
	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get all loot items
	loot_items, err := models.GetAllLoot()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get looted systems for each team
	for _, team := range teams {

		systemIps, err := models.GetLootedTeamSystems(team.Id)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			continue
		}

		// Save to our map
		teamLootedSystems[team.Id] = systemIps
	}

	// Pass to the template
	c.Data["team_looted_systems"] = teamLootedSystems
	c.Data["loot_items"] = loot_items
	c.Data["teams"] = teams
	c.Layout = "sidebar.tpl"
	c.TplName = "loot/browser.html"
	return
}

func (c *LootController) Add() {

	// Make sure we have a file uploaded
	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		c.Ctx.WriteString("Error in file upload: " + err.Error())
		return
	}
	defer file.Close()

	// Get the system ip
	systemIp := c.GetString("system_ip")
	if systemIp == "" {
		c.Ctx.WriteString("Provide a system IP")
		return
	}

	// Fetch the system object
	system, err := models.GetSystem(systemIp)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Check if we have a valid loot tag
	lootTag := c.GetString("loot_tag")
	if !models.LootTagValid(lootTag) {
		allowedTags := strings.Join(models.LootTags, ", ")
		c.Ctx.WriteString("Invalid loot tag. Allowed: " + allowedTags)
		return
	}

	// Retrieve the original filename
	filename := fileHeader.Filename

	// Get the JWT cookie
	jwtCookie, err := c.Ctx.Request.Cookie("session")
	if err != nil {
		c.Ctx.WriteString("Authentication error: " + err.Error())
		return
	}

	// Get the username from the JWT token
	username, err := models.GetCurrentUsername(jwtCookie.Value)
	if err != nil {
		c.Ctx.WriteString("Error extracting username: " + err.Error())
		return
	}

	// Get the random filename
	randomFileName := models.GenerateRandomFilename()

	// Save the loot to a file
	err = c.SaveToFile("file", "uploads/loot/"+randomFileName)
	if err != nil {
		c.Ctx.WriteString("Error saving file: " + err.Error())
		return
	}

	// Assemble the loot object
	loot := models.Loot{
		Tag:      lootTag,
		Uploader: username,
		Name:     filename,
		Filename: randomFileName,
		System:   &system,
	}

	// Add the loot item to the database
	err = models.AddLoot(&loot)
	if err != nil {
		c.Ctx.WriteString("Error adding loot: " + err.Error())
		return
	}

	c.Redirect("/loot", 302) // CHANGE AS NEEDED
	return
}

func (c *LootController) Delete() {

	// Check if we have a loot_id
	lootId, err := c.GetInt("loot_id")
	if (err != nil) || (lootId == 0) {
		c.Ctx.WriteString("Provide a loot id")
		return
	}

	// Delete the loot item
	err = models.DeleteLoot(lootId)
	if err != nil {
		c.Ctx.WriteString("Error adding loot: " + err.Error())
		return
	}

	c.Redirect("/loot", 302) // CHANGE AS NEEDED
	return
}

func (c *LootController) Download() {

	// Parse the loot ID integer
	lootId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if (err != nil) || (lootId == 0) {
		c.Ctx.WriteString("Provide a loot id")
		return
	}

	// Retrieve the file name
	fileName, err := models.GetLootName(lootId)
	if err != nil {
		c.Ctx.WriteString("Failed to get file name: " + err.Error())
		return
	}

	// Sanitize the file name
	safeFileName := govalidator.SafeFileName(fileName)

	// Retrieve the file path
	filePath, err := models.GetLootPath(lootId)
	if err != nil {
		c.Ctx.WriteString("Failed to get file path: " + err.Error())
		return
	}

	// Read the file
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		c.Ctx.WriteString("Failed to read file: " + err.Error())
		return
	}

	// Set response headers
	c.Ctx.Output.Header("Content-Disposition", "attachment; filename=\""+safeFileName+"\"")
	c.Ctx.Output.Header("Content-Type", "application/octet-stream")

	// Return the file contents
	c.Ctx.ResponseWriter.Write(fileContents)
	return
}
