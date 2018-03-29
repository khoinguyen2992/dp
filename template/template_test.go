package template

import (
	"fmt"
)

// ExampleGenerate shows how we can generate 2 messages easily with template pattern
func ExampleGenerate() {
	approveMessage := &ApproveMessage{}
	fmt.Println(Generate(approveMessage))

	rejectMessage := &RejectMessage{}
	fmt.Println(Generate(rejectMessage))
	// Output:
	// Hello, you are approved.
	// Hello, you are rejected.
}
