package main

import (
	"fmt"
	_ "raven/database"
	"raven/models"
	_ "raven/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Get the database name from the config
	dbName, err := beego.AppConfig.String("dbname")
	if err != nil {
		panic(fmt.Errorf("failed to get dbname: %v", err))
	}

	// Start the database
	force, verbose := false, true
	err = orm.RunSyncdb(dbName, force, verbose)
	if err != nil {
		panic(err)
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
	beego.AddFuncMap("formatTime", models.FormatTime)

	beego.Run()
}
