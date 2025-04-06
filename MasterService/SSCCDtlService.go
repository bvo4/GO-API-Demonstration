package MasterService

import (
	"API_DEMONSTRATION/Models"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

/* Source: https://go.dev/doc/tutorial/database-access */
func SSCC_InsertSSCC(ConnectionString Models.SQL_Conn, EpcisDtl Models.ItemsTreeResult) {

	MySQLConn := GetSqlConfig(ConnectionString)
	BulkCopy, err := MySQLConn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := BulkCopy.Prepare(pq.CopyIn("GTIN", "SerialNumber", "LotNum", "Qty"))
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range EpcisDtl.Items {
		_, err = stmt.Exec(user.Gtin, user.SerialNumber, user.LotNum, user.Amount)
		if err != nil {
			log.Fatal(err)
		}
		BulkInsertUsers()
	}

}

func BulkInsertUsers(db *sql.DB, Items []Models.Items) error {
	query := `INSERT INTO users (name, email) VALUES (@userName, @userEmail)`

	batch := &pgx.Batch{}
	for _, user := range Items {
		args := pgx.NamedArgs{
			"Gtin":         user.Gtin,
			"SerialNumber": user.SerialNumber,
			"LotNum":       user.LotNum,
			"Qty":          user.Amount,
		}
		batch.Queue(query, args)
	}

	res, err := dbconn.Conn().PgConn().CopyFrom(context.Background(), f, "COPY csv_test FROM STDIN (FORMAT csv)")
	if err != nil {
		panic(err)
	}
	fmt.Print(res.RowsAffected())
}
