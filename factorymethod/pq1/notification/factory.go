package notification

import "fmt"

func NewNotifier(vertical string) (Notifier, error){
	switch vertical {
	case "sms":
		return &SmsNotification{}, nil
	case "email":
		return &EmailNotification{}, nil
	case "push":
		return &PushNotification{}, nil
	default:
		return nil, fmt.Errorf("unknown notification implementation: %s", vertical)
	}
}