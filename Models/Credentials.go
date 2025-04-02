package Models

type Credentials struct {
	API_KEY   string
	API_VALUE string
}

type Settings struct {
	ConnectionString   string
	_Credentials       Credentials
	EmailServerAddress string
}
