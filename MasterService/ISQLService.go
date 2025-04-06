package MasterService

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// _ is needed for pq to keep it in since it's needed on top of /sql

/* https://go.dev/doc/database/open-handle */
func GetSqlConfig(ConnectionString string) *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		"DESKTOP-91IQ0A5", "", "", 1433)

	db, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal((err))
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("SQL Connection Opened")

	// See "Important settings" section.
	/*
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	*/
	return db
}
