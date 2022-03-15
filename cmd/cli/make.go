package main

import (
	"errors"
	"fmt"
	"time"
)

//doMake() triggers the action of "make" command and their subcommand in CLI
func doMake(arg2, arg3 string) error {

	switch arg2 {
	//make subcommand
	case "migration":
		//posgtres is defaultt dbtype in blacksmith
		dbType := bls.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("yo must give the migration a name"))
		}
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
		upFile := bls.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := bls.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		//create templates for migrations to give the end users of myapp some starting point
		err := copyFilefromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGracefully(err)
		}
		err = copyFilefromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGracefully(err)
		}
	}
	return nil
}
