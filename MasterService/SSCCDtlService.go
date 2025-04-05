package MasterService

import (
	"API_DEMONSTRATION/Models"
)

/* Source: https://go.dev/doc/tutorial/database-access */
func SSCC_InsertSSCC(ConnectionString string, EpcisDtl Models.ItemsTreeResult) {

	//MySQLConn := GetSqlConfig(ConnectionString)
	GetSqlConfig(ConnectionString)
	/*
		_, err := MySQLConn.Begin()
		if err != nil {
			log.Fatal(err)
		}
	*/
	/*
		stmt, err := txn.Prepare(pq.CopyIn("GTIN", "SerialNumber", "LotNum", "Qty"))
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range EpcisDtl.Items {
			_, err = stmt.Exec(user.Gtin, user.SerialNumber, user.LotNum, user.Amount)
			if err != nil {
				log.Fatal(err)
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

		err = txn.Commit()
		if err != nil {
			log.Fatal(err)
		}
	*/
}
