package Models

type Credentials struct {
	API_KEY   string
	API_VALUE string
}

type Settings struct {
	IsDebug     bool
	SQL_Conn    SQL_Conn
	SQL_Target  SQL_Target
	Credentials Credentials
}

type SQL_Conn struct {
	Server   string
	UserId   string
	Password string
	Port     int
}

// Finding the location in the SQL Database
type SQL_Target struct {
	Table   string
	Columns []string
}

func GetSQLTable(SQL_Settings SQL_Target) string {
	return SQL_Settings.Table
}

func GetSQLColumns(SQL_Settings SQL_Target) []string {
	return SQL_Settings.Columns
}
