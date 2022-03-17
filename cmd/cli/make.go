package main

import (
	"errors"
	"fmt"
	"time"
)

//doMake() triggers the action of "make" command and their sub-cmd in CLI
func doMake(arg2, arg3 string) error {
	//arg2 is sub-cmd arg3 is sub-cmd option

	switch arg2 {
	//make subcommand
	case "migration":
		//posgtres is default dbtype in blacksmith
		dbType := bls.DB.DataType
		//sub-cmd option is empty
		if arg3 == "" {
			exitGracefully(errors.New("you must give the migration a name"))
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

	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}
	}
	return nil
}
