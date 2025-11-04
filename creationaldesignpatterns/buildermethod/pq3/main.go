package main

import (
	"fmt"

	"example.com/buildermethod/request"
)

func main() {
	req := request.NewRequestBuilder().
		SetMethod("POST").
		SetURL("https://api.example.com").
		AddHeader("Content-Type", "application/json").
		SetBody("{}").
		Build()
	fmt.Printf("Request generated: %v\n", req)
}
