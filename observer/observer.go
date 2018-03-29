package observer

import (
	"fmt"
)

// We are going to implement a trade platform, in which coins can broadcast updates to their subscribers
// We need components (subscribers) that keeps update by letting other component (coin) trigger its update method
// To keep it simple, we will not implement method to add and remove subscriber of a coin

// Subscriber is an abstract which defines method to trigger by coin
type Subscriber interface {
	Notify(message string)
}

// Coin contains a list of its subscribers who need its update
// Coin also need method to add and remove subscibers, which have not been there yet
type Coin struct {
	subscribers []Subscriber
}

// NotifyTraders is called when coin broadcast update to its subscribers
func (p *Coin) NotifyTraders() {
	for _, s := range p.subscribers {
		s.Notify("gets updated")
	}
}

// Trader ...
type Trader struct {
	Name string
}

// Notify ...
func (f *Trader) Notify(message string) {
	fmt.Println(f.Name, message)
}
