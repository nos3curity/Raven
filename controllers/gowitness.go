package controllers

import (
    "strconv"
    "io/ioutil"
    "log"
    "path/filepath"
    "regexp"
    "raven/models"

    beego "github.com/beego/beego/v2/server/web"
)

type GowitnessController struct {
    beego.Controller
}

// ProcessScreenshots reads screenshot files from a directory, checks for their existence in the database,
// and adds them if they don't exist.
func (c *GowitnessController) ProcessScreenshots() {
    dirPath := "gowitness/screenshots"  // Directory path where screenshots are stored.
    files, err := ioutil.ReadDir(dirPath)  // Read all files in the directory.
    if err != nil {
        log.Fatal(err)  // Log and halt on error while reading the directory.
        return
    }

    // Regex to parse filenames into components: protocol, IP address, and port number.
    regex := regexp.MustCompile(`(http|https)-(\d+\.\d+\.\d+\.\d+)-(\d+)\.png`)
    for _, file := range files {
        matches := regex.FindStringSubmatch(file.Name())  // Match the filename against the regex.
        if len(matches) > 0 {  // If the regex matches, proceed to process the screenshot.
            ip := matches[2]  // Extract the IP address from the filename.
            portNum, _ := strconv.Atoi(matches[3])  // Convert the port number to an integer.
            filePath := filepath.Join(dirPath, file.Name())  // Create the full file path.
            systemPort, err := models.FindSystemPortByIPAndPort(ip, portNum)  // Retrieve the associated SystemPort.
            if err != nil {
                log.Println("Failed to find SystemPort:", err)  // Log if the SystemPort cannot be found.
                continue  // Skip this file and continue with the next one.
            }
            err = models.AddScreenshotToPort(systemPort.Id, filePath)  // Attempt to add the screenshot to the database.
            if err != nil {
                log.Println("Failed to add screenshot:", err)  // Log if the addition fails.
            }
        }
    }
}

// Debuggin
func (c *GowitnessController) LogScreenshotDetails() {
    screenshots, err := models.GetScreenshotDetails() 
    if err != nil {
        log.Println("Error fetching screenshot details:", err)
        c.Ctx.Output.SetStatus(500)
        c.Ctx.Output.Body([]byte("Internal Server Error"))
        return
    }

    for _, screenshot := range screenshots {
        log.Printf("FILENAME: %s associated with IP address: %s on PORT: %d\n",
            screenshot.Filename,
            screenshot.SystemPort.System.Ip,
            screenshot.SystemPort.Port.PortNumber,
        )
    }
    c.Ctx.Output.SetStatus(200)
    c.Ctx.Output.Body([]byte("Screenshot details logged successfully")) 
}
