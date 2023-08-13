package main

import (
	"fmt"
	"sort" // We can use the slices package, but not every Go installation has it.
	"strconv"
	"strings"
)

// This program receives integers from the user, and
// it stores them in a sorted order inside a container.
func main() {
	// declare where we'll store the user's input.
	var input string

	// make the slice of length 0.
	container := make([]int, 0)

	// user input code is placed inside a loop:
	for {
		fmt.Println("\nTo exit, type x.")
		fmt.Println("Please enter an integer to be added to the container: ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "x" {
			return
		}
		if num, err := strconv.Atoi(input); err != nil {
			fmt.Printf("Error: \"%v\" is not an integer (0, 1, 2, 3...).\n", input)
		} else {
			container = append(container, num) // this method is not destructive.
			sort.Ints(container)               // this one is, you see: it changes the slice.
			fmt.Println("\nHere are your numbers, sorted: ")
			fmt.Println("CONTAINER: ", container)
		}
	}
}
