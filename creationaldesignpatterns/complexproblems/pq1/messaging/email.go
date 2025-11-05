package messaging

import (
	"errors"
)

// emailFactory implements PlatformFactory for email.
type emailFactory struct{}

func (f *emailFactory) NewBuilder() Builder {
	// You can set sensible defaults here
	return &emailBuilder{
		headers:      nil,
		footers:      nil,
		inlineImages: nil,
		body:         "",
	}
}

func (f *emailFactory) Service() PlatformService {
	return getEmailServiceInstance()
}

// emailBuilder builds *Message for email platform.
type emailBuilder struct {
	headers      []string
	footers      []string
	inlineImages []string
	body         string
}

func (b *emailBuilder) AddHeader(h string) Builder {
	if h == "" {
		return b
	}
	b.headers = append(b.headers, h)
	return b
}

func (b *emailBuilder) AddFooter(f string) Builder {
	if f == "" {
		return b
	}
	b.footers = append(b.footers, f)
	return b
}

func (b *emailBuilder) AddInlineImage(img string) Builder {
	if img == "" {
		return b
	}
	b.inlineImages = append(b.inlineImages, img)
	return b
}

func (b *emailBuilder) SetBody(body string) Builder {
	b.body = body
	return b
}

// Build validates, applies defaults, and returns a new Message instance (not a singleton).
func (b *emailBuilder) Build() (*Message, error) {
	// Apply defaults
	if b.body == "" {
		// emails should have at least an empty body; could be required in your system
		b.body = "(no body)"
	}

	// validation example: email-specific rule (here simple)
	if len(b.body) > 10000 {
		return nil, errors.New("email body too large")
	}

	msg := &Message{
		headers:      append([]string{}, b.headers...),
		footers:      append([]string{}, b.footers...),
		inlineImages: append([]string{}, b.inlineImages...),
		body:         b.body,
		platform:     "EMAIL",
	}
	return msg, nil
}
