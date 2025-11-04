package logger

import "fmt"

type db struct{}

func (d *db) Log(message string) string {
	return fmt.Sprintf("Logged message: %s in db", message)
}