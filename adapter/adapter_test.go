package adapter

func Example() {
	service := NewService()
	service.Send(Message{
		PhoneNumber: "123456789",
		CountryCode: "84",
		Content:     "content",
	})
	// Output: 84123456789 content
}
