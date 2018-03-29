package strategy

func ExampleLogin() {
	token := "1234"
	fbStrategy := &FacebookStrategy{}
	Login(fbStrategy, token)

	ggStrategy := &GoogleStrategy{}
	Login(ggStrategy, token)

	// Output:
	// facebook 1234
	// google 1234
}
