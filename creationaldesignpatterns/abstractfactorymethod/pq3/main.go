package main

import (
	"fmt"

	"example.com/abstractfactory/cloud"
)

func main() {
	cloudClient, err := cloud.NewCloudFactory("aws")
	if err != nil {
		fmt.Println(err.Error())
	}
	storageSvc := cloudClient.CreateStorageService()
	computeSvc := cloudClient.CreateComputeService()

	fmt.Println(storageSvc.Start())
	fmt.Println(computeSvc.Start())
}
