package blacksmith

import (
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//executes the SQL instruccion to migrations, using go-migrate/migrate CLI integrated
//inside the blacksmith CLI

func (bls *Blacksmith) MigrateUp(dsn string) error {
	//get an open the file to migrations
	m, err := migrate.New("file://"+bls.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	//close the file after to work on it
	defer m.Close()

	//run the migration
	err2 := m.Up()
	if err != nil {
		log.Println("Error running migrations: ", err2)
	}

	return nil
}

func (bls *Blacksmith) MigrateDownAll(dsn string) error {
	//get an open the file to migrations
	m, err := migrate.New("file://"+bls.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	//close the file after to work on it
	defer m.Close()

	err = m.Down()
	if err != nil {
		return err
	}

	return nil

}
func (bls *Blacksmith) Steps(n int, dsn string) error {
	//get an open the file to migrations
	m, err := migrate.New("file://"+bls.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	//close the file after to work on it
	defer m.Close()

	err = m.Steps(n)
	if err != nil {
		return err
	}

	return nil
}

func (bls *Blacksmith) MigrateForce(dsn string) error {
	//get an open the file to migrations
	m, err := migrate.New("file://"+bls.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	//close the file after to work on it
	defer m.Close()

	err = m.Force(-1)
	if err != nil {
		return err
	}

	return nil

}
