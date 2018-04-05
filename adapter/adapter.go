package adapter

import (
	"fmt"
)

// We have a new customer. They want to use our system to send their sms.
// The client wants its own interface which is different from what we currently have.
// We need a layer which meets the new requirements while keeping the current interface the same.

// LegacySender stands for current system which is unabled to modify
type LegacySender struct{}

// Send is function belongs to current system
func (s *LegacySender) Send(phoneNumber, message string) {
	fmt.Println(phoneNumber, message)
}

// Message is a new param we have to receive from new client
type Message struct {
	CountryCode string
	PhoneNumber string
	Content     string
}

// NewSender is middle layer to handle requests from new client
type NewSender struct {
	legacy *LegacySender
}

// Send is function for new client
func (s *NewSender) Send(message Message) {
	phoneNumber := message.CountryCode + message.PhoneNumber
	s.legacy.Send(phoneNumber, message.Content)
}

// NewService returns new sender
func NewService() *NewSender {
	return &NewSender{
		legacy: &LegacySender{},
	}
}
