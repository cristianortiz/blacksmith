package blacksmith

import (
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

//OpenDB open a DB connection with postgres redis or mariaDB engine
func (b *Blacksmith) OpenDB(dbType, dsn string) (*sql.DB, error) {
	if dbType == "postgres" || dbType == "postgresql" {
		dbType = "pgx"
	}

	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, err
	}
	//ping the DB to test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
