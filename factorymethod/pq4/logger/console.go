package logger

import "fmt"

type console struct{}

func (c *console) Log(message string) string {
	return fmt.Sprintf("Logged message: %s in console", message)
}