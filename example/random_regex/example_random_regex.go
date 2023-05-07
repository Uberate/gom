package main

import (
	"fmt"
	"github.com/uberate/gom/pkg/regexp_trans"
)

// Generate the phone number from a phone regex: ^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{4}[*]{4}$
//
// For safe, the last 4 number set to '*'.
func main() {
	// Define regex.
	phoneRegex := "^1(3\\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\\d|9[0-35-9])\\d{4}[*]{4}$"

	// Create a generator, and set now-nanosecond as random seed.
	r := regexp_trans.NewGenerator(regexp_trans.SetSeedByNotTime)

	// Use generator to create value.
	value, err := r.Generate(phoneRegex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Use value anywhere.
	fmt.Println(value)
}
