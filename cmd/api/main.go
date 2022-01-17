package main

import (
	"fmt"

	"server/internal/apihandler"
)

func main() {
	fmt.Println("Hello World.")

	engine := apihandler.CreateEngine()
	engine.Run()
}
