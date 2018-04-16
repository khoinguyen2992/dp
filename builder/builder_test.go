package builder

import (
	"fmt"
)

func Example() {
	director := &Director{}

	citibankProcess := &CitibankProcess{}
	director.SetProcess(citibankProcess)
	director.Build()

	citibankAccount := citibankProcess.GetAccount()
	fmt.Println(citibankAccount.Password, citibankAccount.ReferralCode)

	hsbcProcess := &HSBCProcess{}
	director.SetProcess(hsbcProcess)
	director.Build()

	hsbcAccount := hsbcProcess.GetAccount()
	fmt.Println(hsbcAccount.Password, hsbcAccount.ReferralCode)

	// Output:
	// citi CITI
	// hsbc HSBC
}
