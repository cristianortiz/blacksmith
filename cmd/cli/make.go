package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		//posgtres is defualt dbtype in blacksmith
		dbType := bls.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("yo must give the migration a name"))
		}
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
		upFile := bls.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := bls.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		//create templates for migrations to give the end users of myapp some starting point
	}
	return nil
}
