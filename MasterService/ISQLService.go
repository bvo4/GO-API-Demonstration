package MasterService

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// _ is needed for pq to keep it in since it's needed on top of /sql

/* https://go.dev/doc/database/open-handle */
func GetSqlConfig(ConnectionString string) *sql.DB {

	cfg := mysql.Config{
		User:   "postgres",
		Passwd: "ResearchStationNoc10",
		Net:    "tcp",
		Addr:   "golangapitest.cf8c2qy4a0f6.us-east-2.rds.amazonaws.com,5432",
		DBName: "golangapitest",
	}
	db, err := sql.Open("postgres", cfg.FormatDSN())

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
