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

func (c *LootController) Prepare() {

	// Fetch the sidebar data
	sidebar := &SiderbarController{Controller: c.Controller}
	sidebar.GetTeams()

	teamLootedSystems := make(map[int][]string)
	var lootedTeams []models.Team

	// Get all teams with loot
	teams, err := models.GetLootedTeams()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get looted systems for each team
	for _, team := range teams {

		// Get the team's looted systems
		systemIps, err := models.GetLootedTeamSystems(team)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			continue
		}

		// Retrieve the team object for the sidebar
		lootedTeam, err := models.GetTeam(team)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			continue
		}

		lootedTeams = append(lootedTeams, lootedTeam)

		// Save to our map
		teamLootedSystems[team] = systemIps
	}

	c.Data["loot_tags"] = models.LootTags
	c.Data["looted_teams"] = lootedTeams
	c.Data["team_looted_systems"] = teamLootedSystems
}

func (c *LootController) All() {

	// Get all loot items
	lootItems, err := models.GetAllLoot()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Filter by the loot tag if present
	lootTag := c.GetString("loot_tag")
	if lootTag != "" {
		lootItems, err = models.FilterLootByTag(lootItems, lootTag)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			return
		}

		c.Data["selected_tag"] = lootTag
	}

	// Pass to the template
	c.Data["loot_items"] = lootItems
	c.Layout = "layout/sidebar.tpl"
	c.TplName = "loot/browser.html"
	return
}

func (c *LootController) SystemLoot() {

	// Get the system ip
	systemIp := c.Ctx.Input.Param(":ip")
	if systemIp == "" {
		c.Ctx.WriteString("Provide a system IP")
		return
	}

	// Get the loot for the system
	systemloot, err := models.GetSystemLoot(systemIp)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Filter by the loot tag if present
	lootTag := c.GetString("loot_tag")
	if lootTag != "" {
		systemloot, err = models.FilterLootByTag(systemloot, lootTag)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			return
		}

		c.Data["selected_tag"] = lootTag
	}

	// Pass to the template
	c.Data["loot_items"] = systemloot
	c.Layout = "layout/sidebar.tpl"
	c.TplName = "loot/browser.html"
	return
}

func (c *LootController) TeamLoot() {
	var flattenedLoot []models.Loot

	// Parse the team ID integer
	teamId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Get the IPs of systems looted by the team
	lootedSystems, err := models.GetLootedTeamSystems(teamId)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	// Loop over the looted system IPs and accumulate their loot directly into flattenedLoot
	for _, system := range lootedSystems {
		loot, err := models.GetSystemLoot(system)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			return
		}

		// Append each Loot object to the flattenedLoot slice
		flattenedLoot = append(flattenedLoot, loot...)
	}

	// Filter by the loot tag if present
	lootTag := c.GetString("loot_tag")
	if lootTag != "" {
		flattenedLoot, err = models.FilterLootByTag(flattenedLoot, lootTag)
		if err != nil {
			c.Ctx.WriteString(err.Error())
			return
		}

		c.Data["selected_tag"] = lootTag
	}

	// Pass the flattened loot to the template
	c.Data["loot_items"] = flattenedLoot
	c.Layout = "layout/sidebar.tpl"
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
