package MasterService

import "API_DEMONSTRATION/Models"

func InsertSSCC(ConnectionString string, EpcisDtl Models.ItemsTreeResult) {
	SSCC_InsertSSCC(ConnectionString, EpcisDtl)
}
