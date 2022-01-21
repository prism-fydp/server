package main

import (
	"fmt"

	"server/internal/apihandler"
	"server/internal/db"
)

func main() {
	fmt.Println("Hello World.")

	dbConfig := db.CreateDBConfig(
		"prism_development",
		"prism",
		"prism",
		"localhost",
		"5432",
	)

	engine := apihandler.CreateEngine(dbConfig)
	engine.Run()
}
