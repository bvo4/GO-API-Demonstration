package HTTPServices

import "fmt"

func BuildOrderIdURI(OrderId string) string {

	if len(OrderId) < 1 {
		return ""
	}

	RequestURI := fmt.Sprintf("REDACTED", OrderId)

	return RequestURI

}
