package main

import (
	"fmt"

	"example.com/buildermethod/computer"
)

func main() {
	pc := computer.NewBuilder().
		SetCPU("Intel i9").
		SetRAM("32GB").
		SetStorage("1TB SSD").
		SetGPU("NVIDIA RTX 4080").
		Build()

	fmt.Printf("Computer built: %+v\n", pc)
}
