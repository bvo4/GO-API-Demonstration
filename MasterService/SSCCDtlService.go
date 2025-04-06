package MasterService

import (
	"API_DEMONSTRATION/Models"
	"log"

	mssql "github.com/denisenkom/go-mssqldb"
)

/*
Source: https://go.dev/doc/tutorial/database-access
https://github.com/denisenkom/go-mssqldb/blob/master/examples/bulk/bulk.go
*/
func SSCC_InsertSSCC(ConnectionString Models.SQL_Conn, EpcisDtl Models.ItemsTreeResult) {

	MySQLConn := GetSqlConfig(ConnectionString)
	BulkCopy, err := MySQLConn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	Tx := mssql.CopyIn("[API_DATABASE].[dbo].[SSCC]", mssql.BulkOptions{}, "GTIN", "SerialNumber", "LotNum", "Qty")

	stmt, err := BulkCopy.Prepare(Tx)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range EpcisDtl.Items {
		_, err = stmt.Exec(user.Gtin, user.SerialNumber, user.LotNum, user.Amount)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, user := range EpcisDtl.Containers {
		for _, SubContainer := range user.Items {

			_, err = stmt.Exec(SubContainer.Gtin, SubContainer.SerialNumber, SubContainer.LotNum, SubContainer.Amount)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = BulkCopy.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
