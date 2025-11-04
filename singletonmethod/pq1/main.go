package main

import (
	"fmt"

	"example.com/database/database"
)

func main() {
	db1 := database.GetInstance()
	db2 := database.GetInstance()

	fmt.Println(db1.Query("SELECT * FROM users"))
	fmt.Printf("Same instance: %v\n", db1 == db2)
}
