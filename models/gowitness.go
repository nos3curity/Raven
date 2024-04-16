package models

import (
    "github.com/beego/beego/v2/client/orm"
    "time"
    "log"
)

type Screenshot struct {
    Id         int          `orm:"auto"`
    SystemPort *SystemPort  `orm:"rel(fk);on_delete(cascade)"`
    Filename   string       `orm:"size(255)"`
    Timestamp  time.Time    `orm:"type(datetime);auto_now_add"`
}

func init() {
    orm.RegisterModel(new(Screenshot))
}

// AddScreenshotToPort checks if a screenshot with a given filename and portID already exists.
// If it does not exist, it creates a new Screenshot record in the database.
// This function helps prevent duplicate screenshots for the same port and filename.
func AddScreenshotToPort(portID int, filename string) error {
    o := orm.NewOrm()  // Create a new ORM object to interact with the database.

    // Check if a screenshot with the specified filename and port ID already exists in the database.
    exists := o.QueryTable(new(Screenshot)).Filter("Filename", filename).Filter("SystemPort__Id", portID).Exist()
    if exists {
        log.Println("Screenshot already exists:", filename)  // Log if the screenshot exists to avoid duplication.
        return nil  // Exit without adding a new record.
    }

    // If the screenshot does not exist, create a new Screenshot object.
    screenshot := Screenshot{
        SystemPort: &SystemPort{Id: portID},  // Associate the screenshot with a specific system port.
        Filename: filename,  // Set the filename.
    }
    _, err := o.Insert(&screenshot)  // Insert the new Screenshot into the database.
    if err != nil {
        log.Println("Error inserting screenshot:", err)  // Log any errors that occur during insertion.
        return err  // Return the error to the caller.
    }
    return nil  // Return nil on successful insertion.
}

// GetScreenshotDetails retrieves all screenshots from the database along with their
// related SystemPort details. This function is useful for generating reports or views
// where information about all screenshots is required.
func GetScreenshotDetails() ([]*Screenshot, error) {
    o := orm.NewOrm()  // Create a new ORM object to interact with the database.
    var screenshots []*Screenshot  // Prepare a slice to hold the screenshots.

    // Query the database for all screenshots and load their related SystemPort data.
    _, err := o.QueryTable(new(Screenshot)).RelatedSel("SystemPort").All(&screenshots)
    if err != nil {
        return nil, err  // Return an error if the query fails.
    }
    return screenshots, nil  // Return the list of screenshots.
}
