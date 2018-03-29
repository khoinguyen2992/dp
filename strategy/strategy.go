package strategy

import (
	"fmt"
)

// We are making login module
// The module has to support both facebook and google login
// Both flows have their own way to login
// We wants to have something simple to hide these complex details
// The goal (login) is the same. But there are separate strategies to accomplish

// Strategy is abstract for login strategy, which can be extends to any login strategies we wants
type Strategy interface {
	Exec(token string)
}

// FacebookStrategy ...
type FacebookStrategy struct{}

// Exec ...
func (s *FacebookStrategy) Exec(token string) {
	fmt.Println("facebook", token)
}

// GoogleStrategy ...
type GoogleStrategy struct{}

// Exec ...
func (s *GoogleStrategy) Exec(token string) {
	fmt.Println("google", token)
}

// Login is a bit likely to template method.
// However, it only delegate action to strategy.
func Login(strategy Strategy, token string) {
	strategy.Exec(token)
}
