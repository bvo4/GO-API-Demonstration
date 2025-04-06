package Models

type Credentials struct {
	API_KEY   string
	API_VALUE string
}

type Settings struct {
	IsDebug     bool
	SQL_Conn    SQL_Conn
	Credentials Credentials
}

type SQL_Conn struct {
	Server   string
	UserId   string
	Password string
	Port     int
}
