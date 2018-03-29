package template

import (
	"fmt"
)

// DefaultMessage is common methods
type DefaultMessage struct{}

func (m *DefaultMessage) Greeting() string {
	return "Hello"
}

type ApproveMessage struct {
	DefaultMessage
}

func (m *ApproveMessage) Content() string {
	return "you are approved."
}

type RejectMessage struct {
	DefaultMessage
}

func (m *RejectMessage) Content() string {
	return "you are rejected."
}

// ExampleGenerate shows how we can generate 2 messages easily with template pattern
func ExampleGenerate() {
	fmt.Println(Generate(&ApproveMessage{}))
	fmt.Println(Generate(&RejectMessage{}))
	// Output:
	// Hello, you are approved.
	// Hello, you are rejected.
}
