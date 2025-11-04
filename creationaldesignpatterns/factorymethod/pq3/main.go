package main

import (
	"fmt"

	"example.com/factorymethod/gateway"
)

func main() {
	gateway, err := gateway.NewPaymentGateway("paypal")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	fmt.Println(gateway.Process(69))
}