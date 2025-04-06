package MasterService

import (
	"API_DEMONSTRATION/Models"
	"database/sql"

	mssql "github.com/denisenkom/go-mssqldb"
)

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
