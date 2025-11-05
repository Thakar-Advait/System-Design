package messaging

import (
	"errors"
)

// smsFactory implements PlatformFactory for SMS.
type smsFactory struct{}

func (f *smsFactory) NewBuilder() Builder {
	return &smsBuilder{}
}

func (f *smsFactory) Service() PlatformService {
	return getSMSServiceInstance()
}

// smsBuilder builds *Message for SMS platform.
type smsBuilder struct {
	headers      []string
	footers      []string
	inlineImages []string // SMS might not support images but we keep builder interface consistent
	body         string
}

func (b *smsBuilder) AddHeader(h string) Builder {
	if h == "" {
		return b
	}
	b.headers = append(b.headers, h)
	return b
}

func (b *smsBuilder) AddFooter(f string) Builder {
	if f == "" {
		return b
	}
	b.footers = append(b.footers, f)
	return b
}

func (b *smsBuilder) AddInlineImage(img string) Builder {
	// Some SMS providers support MMS — allow adding but it's optional
	if img == "" {
		return b
	}
	b.inlineImages = append(b.inlineImages, img)
	return b
}

func (b *smsBuilder) SetBody(body string) Builder {
	b.body = body
	return b
}

func (b *smsBuilder) Build() (*Message, error) {
	// defaults & validation; SMS often has length constraints
	if b.body == "" {
		return nil, errors.New("sms body is required")
	}
	if len(b.body) > 1600 {
		return nil, errors.New("sms body too long")
	}
	msg := &Message{
		headers:      append([]string{}, b.headers...),
		footers:      append([]string{}, b.footers...),
		inlineImages: append([]string{}, b.inlineImages...),
		body:         b.body,
		platform:     "SMS",
	}
	return msg, nil
}
