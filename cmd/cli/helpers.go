package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

//setup() populates some values in blacksmith type specific for migrations
func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}
	//get working directory, also path will be rootPath
	path, err := os.Getwd()
	if err != nil {
		exitGracefully(err)
	}

	bls.RootPath = path
	bls.DB.DataType = os.Getenv("DATABASE_TYPE")
}

//getDSN( ) returns the DSN in the correct format depending
//on wich dbType migrations are running
func getDSN() string {
	dbType := bls.DB.DataType
	//if DB is defined as pgx set postgres
	if dbType == "pgx" {
		dbType = "postgres"
	}
	//config dsn for postgres
	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))

		}
		return dsn
	}
	//config dsn for mysql
	return "mysql://" + bls.BuildDSN()

}

//showHelp() print a text with help info about the cli commands and their functions
func showHelp() {
	color.Yellow(`Available commands:

	help        			 - show the help commands
	version     		 	 - print app version
	migrate     		 	 - runs all up migrations that have not been run previously
	migrate down 		 	 - reverses the most recent migration
	migration reset		 	 - runs all down migrations in reverse order, and then all up migrations
	make migrations <name>   - creates two new up and down migrations in the migrations folder
	make auth                - creates and runs migrations for authentication tables, and creates models and middleware
	make handler <name>      - creates a stub handler file in the handler directory
	make model <model>		 - creates new model file in the data directory
	make session			 - creates a table in DB as session store
	`)
}
