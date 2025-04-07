package MasterService

import (
	"API_DEMONSTRATION/Models"
	"database/sql"
	"fmt"

	mssql "github.com/denisenkom/go-mssqldb"
)

/* Prepares the SQL Transaction statement and copies the item set in */
func SQLPrepareBulkCopy(SQL_SETTINGS Models.SQL_Target, BulkCopy *sql.Tx, EpcisDtl []Models.Items) *sql.Stmt {
	Tx := mssql.CopyIn(SQL_SETTINGS.Table, mssql.BulkOptions{}, SQL_SETTINGS.Columns...)

	stmt, err := BulkCopy.Prepare(Tx)
	Models.CheckError(err)

	/* Loads the data into the SQL transaction */
	for _, OrderItem := range EpcisDtl {
		_, err = stmt.Exec(OrderItem.Gtin, OrderItem.SerialNumber, OrderItem.LotNum, OrderItem.Amount)
		Models.CheckError(err)
	}
	return stmt
}

/* Initiates the SQL Transaction */
func SQLExecStatement(stmt *sql.Stmt, BulkCopy *sql.Tx) {
	_, err := stmt.Exec()
	Models.CheckError(err)

	err = stmt.Close()
	Models.CheckError(err)

	err = BulkCopy.Commit()
	Models.CheckError(err)

	fmt.Println("Transaction Successful")
}
