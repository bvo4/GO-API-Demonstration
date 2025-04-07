package MasterService

import (
	"API_DEMONSTRATION/Models"
)

/*
Source: https://go.dev/doc/tutorial/database-access
https://github.com/denisenkom/go-mssqldb/blob/master/examples/bulk/bulk.go
*/
func SSCC_InsertSSCC(CREDENTIALS Models.Settings, EpcisDtl []Models.Items) {

	MySQLConn := GetSqlConfig(CREDENTIALS.SQL_Conn)
	BulkCopy, err := MySQLConn.Begin()
	Models.CheckError(err)

	//Prepare the batch with the bulk
	stmt := SQLPrepareBulkCopy(CREDENTIALS.SQL_Target, BulkCopy, EpcisDtl)

	//Execute Sql Transaction
	SQLExecStatement(stmt, BulkCopy)

}
