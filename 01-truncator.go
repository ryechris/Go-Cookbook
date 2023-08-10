package main

import (
	"fmt"
	"strconv"
)

// program removes all the numbers to the right of the decimal point.
func main() {
	var entry float64 // declare the var to store userinput in.
	fmt.Printf("Enter a float64: (e.g. 3.1415): ") // prompt user for input

	_, e := fmt.Scan(&entry) // store it in where that pointer is pointing at.

	if e != nil {
		fmt.Println("Please enter a floating number, such as")
	}

	entry_truncated := strconv.FormatFloat(entry, 'f', 0, 64) // convert the float to string.
	fmt.Printf("The number is now a string & truncated: %s\n", entry_truncated) // print the result

}

