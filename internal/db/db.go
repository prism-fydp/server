package db

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func validateSchema(dbConfig *DBConfig) error {
	absPath, err := filepath.Abs("migrations")
	if err != nil {
		return err
	}

	m, err := migrate.New(
		fmt.Sprintf("%s%s", "file:///", absPath),
		createDatabaseURL(dbConfig)+"?sslmode=disable",
	)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func CreateDBClient(dbConfig *DBConfig) *pgx.Conn {
	// Connect to the database
	conn, err := pgx.Connect(context.Background(), createDatabaseURL(dbConfig))
	if err != nil {
		log.Fatal("Error: could not connect to database.")
	}

	err = validateSchema(dbConfig)
	if err != nil {
		log.Print("Error: ", err)
		log.Fatal("Error: could not migrate database")
	}

	return conn
}
