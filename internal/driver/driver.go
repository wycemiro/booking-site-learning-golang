package driver

import (
	"database/sql"
	"time"
)

//DB holds database conn pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10

const maxIdleDbConn = 5

const maxDbLifetime = 5 * time.Minute

//ConnectSQL creates db pool for postgresql
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)
	dbConn.SQL = d
	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

//testDB tries to ping db
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

//NewDatabase creates a new db for the app
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}
	//check if db OK by ping
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}