package HTTPServices

func BuildOrderIdURI(OrderId string) string {

	if len(OrderId) < 1 {
		return ""
	}

	return RequestURI

}
