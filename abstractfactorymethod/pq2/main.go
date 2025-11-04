package main

import (
	"fmt"

	"example.com/abstractfactory/database"
)

func main() {
	// Create MySQL factory
	mysqlFactory, err := database.NewDBFactory("mysql")
	if err != nil {
		panic(err)
	}

	// Create MySQL products
	mysqlConn := mysqlFactory.CreateConnection()
	mysqlQB := mysqlFactory.CreateQueryBuilder()

	fmt.Println(mysqlConn.Connect())
	fmt.Println(mysqlQB.Build())

	// Create Postgres factory
	postgresFactory, err := database.NewDBFactory("postgres")
	if err != nil {
		panic(err)
	}

	// Create Postgres products
	postgresConn := postgresFactory.CreateConnection()
	postgresQB := postgresFactory.CreateQueryBuilder()

	fmt.Println(postgresConn.Connect())
	fmt.Println(postgresQB.Build())
}
