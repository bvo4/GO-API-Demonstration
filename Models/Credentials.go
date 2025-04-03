package Models

type Credentials struct {
	API_KEY   string
	API_VALUE string
}

type Settings struct {
	IsDebug          bool
	ConnectionString string
	Credentials      Credentials
}
