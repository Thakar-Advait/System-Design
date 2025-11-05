package messaging

import "strings"

// Message is the product built by Builder.
// Fields are unexported to enforce immutability from outside this package.
type Message struct {
	headers      []string
	footers      []string
	inlineImages []string
	body         string
	platform     string
}

// getters (exported read-only accessors)
func (m *Message) Headers() []string      { return append([]string{}, m.headers...) }
func (m *Message) Footers() []string      { return append([]string{}, m.footers...) }
func (m *Message) InlineImages() []string { return append([]string{}, m.inlineImages...) }
func (m *Message) Body() string           { return m.body }
func (m *Message) Platform() string       { return m.platform }

// Helper for debugging/logging
func (m *Message) Summary() string {
	return "Platform: " + m.platform +
		" | Headers: " + strings.Join(m.headers, ",") +
		" | Footers: " + strings.Join(m.footers, ",") +
		" | Images: " + strings.Join(m.inlineImages, ",") +
		" | Body: " + m.body
}
