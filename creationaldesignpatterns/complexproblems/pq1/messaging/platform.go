package messaging

import "time"

// PlatformService represents the delivery service for a platform.
// Implementations (EmailService, SMSService) are singletons.
type PlatformService interface {
	Send(m *Message) error
	Schedule(m *Message, when time.Time) error
	PlatformName() string
}

// Builder builds messages for a platform.
type Builder interface {
	AddHeader(h string) Builder
	AddFooter(f string) Builder
	AddInlineImage(img string) Builder
	SetBody(body string) Builder
	Build() (*Message, error)
}

// PlatformFactory returns platform-specific builder and service.
type PlatformFactory interface {
	NewBuilder() Builder
	Service() PlatformService
}
