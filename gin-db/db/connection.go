package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbConn *sqlx.DB

// Init database connection
func Init(dataSource string) error {
	var err error
	dbConn, err = sqlx.Connect("postgres", dataSource)
	return err
}

// Get shared db connection
func GetConn() *sqlx.DB {
	if dbConn == nil {
		panic("db connection is nil")
	}

	return dbConn
}
