package notification

import "fmt"

type EmailNotification struct{}

func (en *EmailNotification) Send(message string) string {
	return fmt.Sprintf("Sending email notification: %s", message)
}