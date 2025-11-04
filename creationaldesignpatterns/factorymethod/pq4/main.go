package main

import (
	"fmt"

	"example.com/factorymethod/logger"
)

func main() {
	logger, err := logger.NewLogger("db")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(logger.Log("Hello World!"))
}
