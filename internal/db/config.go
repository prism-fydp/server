package db

type DBConfig struct {
	dbname   string
	username string
	password string
	host     string
	port     string
}

func CreateDBConfig(
	dbname string,
	username string,
	password string,
	host string,
	port string,
) *DBConfig {
	dbconfig := DBConfig{
		dbname:   dbname,
		username: username,
		password: password,
		host:     host,
		port:     port,
	}
	return &dbconfig
}
