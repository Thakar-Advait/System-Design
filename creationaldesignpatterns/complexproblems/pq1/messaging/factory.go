package messaging

import "fmt"

// NewPlatformFactory returns a PlatformFactory for a given vertical.
// Add new platforms here (WhatsApp, Push, etc.) easily.
func NewPlatformFactory(vertical string) (PlatformFactory, error) {
	switch vertical {
	case "email":
		return &emailFactory{}, nil
	case "sms":
		return &smsFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported platform: %s", vertical)
	}
}
