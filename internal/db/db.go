package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func createDatabaseURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbConfig.username,
		dbConfig.password,
		dbConfig.host,
		dbConfig.port,
		dbConfig.dbname,
	)
}

func CreateDBClient(dbConfig *DBConfig) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), createDatabaseURL(dbConfig))
	if err != nil {
		// TODO: setup logging
		fmt.Println("Error: could not connect to database.")
		os.Exit(1)
	}

	return conn
}
