package MasterService

import (
	"API_DEMONSTRATION/Models"
	"database/sql/driver"
	"log"

	"github.com/go-sql-driver/mysql"
)

/* Source: https://go.dev/doc/tutorial/database-access */
func SSCC_InsertSSCC(ConnectionString string, EpcisDtl Models.ItemsTreeResult) {

	MySQLConn := GetSqlConfig()

	conn, _ := MySQLConn.Conn(*MySQLConn)
	conn.Raw(func(conn any) error {
		ex := conn.(driver.Execer)
		res, err := ex.Exec(`[dbo].[InsertSSCC]`, nil)
		// Both slices have 2 elements.
		log.Print(res.(mysql.Result).AllRowsAffected())
		log.Print(res.(mysql.Result).AllLastInsertIds())
	})
}
