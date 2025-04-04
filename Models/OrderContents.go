package Models

type OrderContents struct {
	Order_ID     string
	GTIN         string
	SerialNumber string
	LotNum       string
}

func ToOrderContents(Records [][]string, i int) OrderContents {
	var Temp OrderContents
	Temp.Order_ID = Records[i][0]
	Temp.GTIN = Records[i][1]
	Temp.SerialNumber = Records[i][2]
	Temp.LotNum = Records[i][3]
	return Temp
}
