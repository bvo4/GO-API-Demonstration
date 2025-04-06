package MasterService

import (
	"API_DEMONSTRATION/Models"
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
