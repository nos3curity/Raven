package database

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	runmode, err := beego.AppConfig.String("runmode")
	if err != nil {
		panic(fmt.Errorf("failed to get run mode: %v", err))
	}

	var dataSource string

	// Get the database name from the config
	dbName, err := beego.AppConfig.String("dbname")
	if err != nil {
		panic(fmt.Errorf("failed to get dbname: %v", err))
	}

	if runmode == "dev" {

		fmt.Println("Running in dev mode. Using SQLite.")

		// Get the datasource for the dev environment
		dataSource, err = beego.AppConfig.String("dev::datasource")
		if err != nil {
			panic(fmt.Errorf("failed to get dev datasource: %v", err))
		}

		// Register the database driver
		err := orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if err != nil {
			panic(fmt.Errorf("failed to register the sqlite driver: %v", err))
		}

		// Create the database
		err = orm.RegisterDataBase(dbName, "sqlite3", dataSource)
		if err != nil {
			panic(fmt.Errorf("failed to register the default database: %v", err))
		}

	} else if runmode == "prod" {

		fmt.Println("Running in prod mode. Using PostgreSQL.")

		// Get the datasource for the prod environment
		dataSource, err = beego.AppConfig.String("prod::datasource")
		if err != nil {
			panic(fmt.Errorf("failed to get prod datasource: %v", err))
		}

		// PostgreSQL registration logic here
		err = orm.RegisterDriver("postgres", orm.DRPostgres)
		if err != nil {
			panic(fmt.Errorf("failed to register the postgres driver: %v", err))
		}

		err = orm.RegisterDataBase(dbName, "postgres", dataSource)
		if err != nil {
			panic(fmt.Errorf("failed to register the postgres database: %v", err))
		}
	}

	orm.DefaultTimeLoc = time.FixedZone("UTC-8", -8*60*60)
}
