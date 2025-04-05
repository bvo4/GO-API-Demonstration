package MasterService

import (
	"database/sql"
	"log"
	"time"
)

/* https://go.dev/doc/database/open-handle */
func GetSqlConfig(ConnectionString string) *sql.DB {

	db, err := sql.Open("mysql", ConnectionString)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
