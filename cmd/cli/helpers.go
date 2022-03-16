package main

import (
	"fmt"
	"os"

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
	//if DB is pgx set postgres
	if dbType == "pgx" {
		dbType = "postgres"
	}
	//config dsn for posgress
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
