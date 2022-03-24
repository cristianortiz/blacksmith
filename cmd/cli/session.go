package main

import (
	"fmt"
	"time"
)

// doSessionTable() executes the code of make session sub-cmd, to create the up and down
// migrations to create session table in DB
func doSessionTable() error {
	//first need to know what DB is being used
	dbType := bls.DB.DataType

	if dbType == "mariadb" {
		dbType = "mysql"

	}

	if dbType == "postgresql" {
		dbType = "postgres"

	}
	//migration file for session table, located in myapp/migrations
	fileName := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro())
	upFile := bls.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := bls.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

	//copy the sql up migration from the migrations template folder
	err := copyFilefromTemplate("templates/migrations/"+dbType+"_session.sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	//to downfile just copy in it a simple drop table sql sentence
	err = copyDataToFile([]byte("drop table sessions"), downFile)
	if err != nil {
		exitGracefully(err)
	}
	//execute the up migration
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	return nil
}
