package main

import (
	"fmt"

	"example.com/singletonmethod/logger"
)

func main() {
	logger1 := logger.GetLoggerInstance()
	logger2 := logger.GetLoggerInstance()

	fmt.Println(logger1.Log("New message"))
	fmt.Printf("Same instance: %v\n", logger1 == logger2)
}
