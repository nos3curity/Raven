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

func AddScreenshotToPort(portID int, filename string) error {
    o := orm.NewOrm()
    screenshot := Screenshot{
        SystemPort: &SystemPort{Id: portID}, // Ensure portID is correct
        Filename: filename,
    }
    _, err := o.Insert(&screenshot)
    if err != nil {
        log.Println("Error inserting screenshot:", err)
        return err
    }
    return nil
}

// GetScreenshotDetails retrieves screenshots with their corresponding system IP and port number.
func GetScreenshotDetails() ([]*Screenshot, error) {
    o := orm.NewOrm()
    var screenshots []*Screenshot
    _, err := o.QueryTable(new(Screenshot)).RelatedSel("SystemPort").All(&screenshots)
    if err != nil {
        return nil, err
    }
    return screenshots, nil
}

