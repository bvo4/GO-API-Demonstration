package MasterService

import (
	"API_DEMONSTRATION/Models"
	"database/sql"

	mssql "github.com/denisenkom/go-mssqldb"
)

/*
Source: https://go.dev/doc/tutorial/database-access
https://github.com/denisenkom/go-mssqldb/blob/master/examples/bulk/bulk.go
*/
func SSCC_InsertSSCC(ConnectionString Models.SQL_Conn, EpcisDtl []Models.Items) {

	MySQLConn := GetSqlConfig(ConnectionString)
	BulkCopy, err := MySQLConn.Begin()
	Models.CheckError(err)

	//Prepare the batch with the bulk
	stmt := SQLPrepareBulkCopy(BulkCopy, EpcisDtl)

	//Execute Sql Transaction
	SQLExecStatement(stmt, BulkCopy)
}

/* Prepares the SQL Transaction statement and copies the item set in */
func SQLPrepareBulkCopy(BulkCopy *sql.Tx, EpcisDtl []Models.Items) *sql.Stmt {
	Tx := mssql.CopyIn("[API_DATABASE].[dbo].[SSCC]", mssql.BulkOptions{}, "GTIN", "SerialNumber", "LotNum", "Qty")

	stmt, err := BulkCopy.Prepare(Tx)
	Models.CheckError(err)

	for _, user := range EpcisDtl {
		_, err = stmt.Exec(user.Gtin, user.SerialNumber, user.LotNum, user.Amount)
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
}
