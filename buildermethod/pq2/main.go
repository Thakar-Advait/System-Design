package main

import (
	"fmt"

	"example.com/buildermethod/pizza"
)

func main() {
	pizza := pizza.NewBuilder().
		SetBase("Thin Crust").
		SetSize("Large").
		AddTopping("Mushrooms").
		AddTopping("Olives").
		Build()

	fmt.Printf("Pizza: %v\n", pizza)
}
