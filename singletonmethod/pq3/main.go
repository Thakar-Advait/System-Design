package main

import (
	"fmt"

	"example.com/singletonmethod/config"
)

func main() {
	configIns1 := config.GetConfigInstace()
	configIns2 := config.GetConfigInstace()
	fmt.Println(configIns1.Get("KEY"))
	fmt.Printf("Same instance: %v\n", configIns1 == configIns2)
}
