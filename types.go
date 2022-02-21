package blacksmith

import "database/sql"

//initPath is a struct to define the main path and folders structure for the WebApp
type initPaths struct {
	//the working directory path to create the webApp folder structure
	rootPath    string
	folderNames []string
}

//cookieConfig type to configure sessionManager
type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}

//databaseConfig type to config private db type and string connection por blacksmith
type databaseConfig struct {
	dsn      string
	database string
}

//Database type to config specific DB engine name and connection pool to myapp
type Database struct {
	DataType string
	Pool     *sql.DB
}
