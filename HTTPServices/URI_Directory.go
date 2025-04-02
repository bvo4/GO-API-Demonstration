package HTTPServices

import "fmt"

func BuildOrderIdURI(OrderId string) string {

	if len(OrderId) < 1 {
		return ""
	}

	RequestURI := fmt.Sprintf("https://onescan.lspedia.com/ext/api/orders/%s/itemstree", OrderId)
	return RequestURI
}
