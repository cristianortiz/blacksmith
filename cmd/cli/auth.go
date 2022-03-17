package main

import (
	"fmt"
	"log"
	"time"
)

//doAuth executes the auth cmd, runs the migrations to create the necessary
//authentication tables in DB
func doAuth() error {
	//migrations  for authentication works
	dbType := bls.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := bls.RootPath + "/migrations" + fileName + ".up.sql"
	downFile := bls.RootPath + "/migrations" + fileName + ".down.sql"
	log.Println(dbType, upFile, downFile)

	//copy the sql template to create auth tables into migration up file
	err := copyFilefromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}
	//copy the sql sentence into  migration down file
	err = copyDataToFile([]byte("drop table if exists users cascade"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	//run auth migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}
	//copy files over
	return nil
}
