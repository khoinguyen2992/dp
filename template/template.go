package template

import (
	"fmt"
)

// We are going to build a notification module.
// It needs a way to generate message included title and content in many use cases
// But message should be in unified format. Title will be followed by content.
// The way to do thing (message generation/ algorithm) is the same.
// But actual parts of it should be open to modified across use cases

// Message is abstract for message, which can be extended to any messages we want
type Message interface {
	Greeting() string
	Content() string
}

// Generate is template method to render message following a fixed format
func Generate(m Message) string {
	return fmt.Sprintf("%s, %s", m.Greeting(), m.Content())
}

// DefaultMessage is common methods
type DefaultMessage struct{}

// Greeting ...
func (m *DefaultMessage) Greeting() string {
	return "Hello"
}

// ApproveMessage ...
type ApproveMessage struct {
	DefaultMessage
}

// Content ...
func (m *ApproveMessage) Content() string {
	return "you are approved."
}

// RejectMessage ...
type RejectMessage struct {
	DefaultMessage
}

// Content ...
func (m *RejectMessage) Content() string {
	return "you are rejected."
}
