package Models

type ItemsTreeResult struct {
	Items      []Items      `json:"items"`
	Containers []Containers `json:"containers"`
}

type Items struct {
	Upc          string `json:"upc"`
	Ndc          string `json:"ndc"`
	Gtin         string `json:"gtin"`
	SerialNumber string `json:"serialNumber"`
	LotNum       string `json:"lotNum"`
	Expiry       string `json:"expiry"`
	Amount       int    `json:"amount"`
}

type Containers struct {
	LotNum       string       `json:"lotNum"`
	Expiry       string       `json:"expiry"`
	SerialNumber string       `json:"serialNumber"`
	Role         int          `json:"role"`
	Number       string       `json:"number"`
	Items        []Items      `json:"items"`
	Containers   []Containers `json:"containers"`
}
