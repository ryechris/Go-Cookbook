package main

import (
	"fmt"
	"strings"
)

// This progarm prompts the user to enter a string.
// It then searches the string to see if the input contains the
// three letters i b m, in that order.
func main() {
	// Declare the vars we need.
	var i, b, m int
	var input string

	// Prompt the user for input.
	fmt.Println("Enter a string (e.g. airbnm, asdf, icbm, etc): ")
	_, e := fmt.Scan(&input)

	if e != nil {
		fmt.Printf("An error occurred: %s", e)
		return
	}

	// We need to make sure the program succeeds whatever case the user gives.
	input = strings.ToLower(input)

	// check the indices for i b m
	i = strings.Index(input, "i")
	b = strings.Index(input, "b")
	m = strings.Index(input, "m")

	switch {
	case i >= 0 && b > i && m > b:
		// input.sliceinto its constituent letters
		s := strings.Split(input, "")

		// search the array and turn it into upper case
		for index := 0; index < len(s); index++ {
			if s[index] == "i" || s[index] == "b" || s[index] == "m" {
				s[index] = strings.ToUpper(s[index]) // ToUpper is nondestructive
			}
		}

		// join the array into one string named capitalized
		capitalized := strings.Join(s, "")

		// return capitalized to user.
		fmt.Printf("Found: %s!\n", capitalized)

	default:
		fmt.Println("The string I B M not found.")
	} // end of switch statement
}
