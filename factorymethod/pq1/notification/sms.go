package notification

import "fmt"

type SmsNotification struct{}

func (sn *SmsNotification) Send(message string) string {
	return fmt.Sprintf("Sending Sms Notification: %s", message)
}