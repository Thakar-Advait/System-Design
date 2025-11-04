package logger

import "fmt"

type file struct{}

func (f *file) Log(message string) string {
	return fmt.Sprintf("Logged message: %s in file", message)
}