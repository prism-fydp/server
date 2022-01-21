package main

import (
	"fmt"

	"server/internal/apihandler"
	"server/internal/db"
)

func main() {
	fmt.Println("Hello World.")

	dbConfig := db.CreateDBConfig(
		"test",
		"test_user",
		"test_password",
		"localhost",
		"5432",
	)

	engine := apihandler.CreateEngine(dbConfig)
	engine.Run()
}
