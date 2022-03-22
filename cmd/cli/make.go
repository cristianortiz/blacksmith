package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

//doMake() triggers the action of "make" command and their sub-cmd in CLI
func doMake(arg2, arg3 string) error {
	//arg2 is sub-cmd arg3 is sub-cmd option

	switch arg2 {
	//make migration sub-command
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
		//make auth sub-cmd
	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}
		//make handler sub-cmd
	case "handler":
		//make handler 'name' as arg3
		if arg3 == "" {
			exitGracefully(errors.New("you must give the handler a name"))
		}
		//build file name for the new handler in myapp/handler folder
		fileName := bls.RootPath + "/handlers/" + strings.ToLower(arg3) + ".go"
		if fileExists(fileName) {
			exitGracefully(errors.New(fileName + " already exists!"))

		}
		//copying template into handler file
		data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
		if err != nil {
			exitGracefully(err)
		}
		handler := string(data)
		//the template handler func() name must be replaced in code it self with the name
		//entered in CLI through: make handler 'handlername'(arg3) command
		handler = strings.ReplaceAll(handler, "$HANDLERNAME$", strcase.ToCamel(arg3))
		err = copyDataToFile([]byte(handler), fileName)
		if err != nil {
			exitGracefully(err)
		}

	case "model":
		//make model 'name' as arg3
		if arg3 == "" {
			exitGracefully(errors.New("you must give the model a name"))
		}

		//copying template into model file
		data, err := templateFS.ReadFile("templates/data/model.go.txt")
		if err != nil {
			exitGracefully(err)
		}
		model := string(data)
		// 'pluralize' the name of the model with a dedicated go package
		plur := pluralize.NewClient()
		//  in model file code  the modelName structs or vars must be in singular and DB tableName vars  must be plurar
		// following the previous naming convention in myapp/data package
		var modelName = arg3
		var tableName = arg3
		//check if the model name entered in command is already a plurar
		if plur.IsPlural(arg3) {
			modelName = plur.Singular(arg3)
			tableName = strings.ToLower(tableName)
		} else {
			tableName = strings.ToLower(plur.Plural(arg3))

		}
		//build file name for the new model in myapp/data folder
		fileName := bls.RootPath + "/data/" + strings.ToLower(modelName) + ".go"
		if fileExists(fileName) {
			exitGracefully(errors.New(fileName + " already exists!"))

		}
		model = strings.ReplaceAll(model, "$MODELNAME$", strcase.ToCamel(modelName))
		model = strings.ReplaceAll(model, "$TABLENAME$", tableName)
		err = copyDataToFile([]byte(model), fileName)
		if err != nil {
			exitGracefully(err)
		}

	}
	return nil
}
