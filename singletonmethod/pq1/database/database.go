package database

import (
	"fmt"
	"sync"
)

type Database struct {
	connection string
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
			fmt.Println("Creating a new database connection...")
			instance = &Database{connection: "Connected to MongoDB"}
		})
	return instance
}

func (d *Database) Query(sql string) string {
	return fmt.Sprintf("Executing: %s on %s", sql, d.connection)
}
