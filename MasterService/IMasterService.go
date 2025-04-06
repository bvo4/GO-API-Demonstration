package MasterService

import "API_DEMONSTRATION/Models"

func InsertSSCC(ConnectionString Models.SQL_Conn, EpcisDtl Models.ItemsTreeResult) {
	SSCC_InsertSSCC(ConnectionString, EpcisDtl)
}
