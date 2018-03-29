package observer

func Example() {
	bob := &Trader{Name: "Bob"}
	alice := &Trader{Name: "Alice"}

	eth := &Coin{
		subscribers: []Subscriber{bob, alice},
	}

	eth.NotifyTraders()
	// Output:
	// Bob gets updated
	// Alice gets updated
}
