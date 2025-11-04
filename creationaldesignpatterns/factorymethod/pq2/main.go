package main

import (
	"fmt"
	"os"

	"example.com/factorymethod/exporter"
)

func main() {
	exporter, err := exporter.NewExporter("HTML")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	dir, _ := os.Getwd()
	fmt.Println(exporter.Export(dir))
}
