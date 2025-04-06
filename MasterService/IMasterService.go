package MasterService

import "API_DEMONSTRATION/Models"

func InsertSSCC(CREDENTIALS Models.Settings, EpcisDtl []Models.Items) {
	SSCC_InsertSSCC(CREDENTIALS, EpcisDtl)
}
