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

func (c *GowitnessController) ProcessScreenshots() {
    dirPath := "gowitness/screenshots"
    files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        log.Fatal(err)
        return
    }

    regex := regexp.MustCompile(`(http|https)-(\d+\.\d+\.\d+\.\d+)-(\d+)\.png`)
    for _, file := range files {
        matches := regex.FindStringSubmatch(file.Name())
        if len(matches) > 0 {
            ip := matches[2]
            portNum, _ := strconv.Atoi(matches[3]) // Renamed to portNum to differentiate from SystemPort object
            filePath := filepath.Join(dirPath, file.Name())
            systemPort, err := models.FindSystemPortByIPAndPort(ip, portNum) // Correct variable and error handling
            if err != nil {
                log.Println("Failed to find SystemPort:", err)
                continue
            }
            err = models.AddScreenshotToPort(systemPort.Id, filePath) // Use systemPort.ID
            if err != nil {
                log.Println("Failed to add screenshot:", err)
            }
        }
    }
}


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
