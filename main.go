package main

import (
	"fmt"
	"raven/models"
	_ "raven/routers"
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:data.db?cache=shared&mode=rwc")
}

func formatTime(t time.Time) string {

	localTime := t.Local()
	return localTime.Format("2006-01-02 15:04:05")
}

func main() {

	// Start the database
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize the JWT cookie signing key
	jwtSecret, _ := models.GetConfig("jwt_secret")
	if jwtSecret.Value == "" {
		models.InitializeJwtSecret()
	}

	// Initialize the server password
	password, _ := models.GetConfig("password")
	if password.Value == "" {
		err := models.InitializePassword()
		password, _ = models.GetConfig("password")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Server Password:", password.Value)

	// Register time template function
	beego.AddFuncMap("formatTime", formatTime)

	beego.Run()
}
