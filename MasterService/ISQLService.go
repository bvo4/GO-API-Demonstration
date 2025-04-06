package MasterService

import (
	"API_DEMONSTRATION/Models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// _ is needed for pq to keep it in since it's needed on top of /sql

/* https://go.dev/doc/database/open-handle */
func GetSqlConfig(ConnectionString Models.SQL_Conn) *sql.DB {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		ConnectionString.Server, ConnectionString.UserId, ConnectionString.Password, ConnectionString.Port)

	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal((err))
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("SQL Connection Opened")

	// See "Important settings" section.
	/*
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	*/
	return db
}
