package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver mysql
)

// Connection open database connection
func Connection() (*sql.DB, error) {
	stringConnection := "golang:golang@/goCourse?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
