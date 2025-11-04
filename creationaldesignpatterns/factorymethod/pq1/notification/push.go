package notification

import "fmt"

type PushNotification struct{}

func (pn *PushNotification) Send(message string) string {
	return fmt.Sprintf("Sending Push Notification: %s", message)
}