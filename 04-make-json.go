package main

import (
	"encoding/json"
	"fmt"
)

// Program takes data you input,
// stores it in one of Go's data types,
// and converts it to JSON to print back to you.
func main() {
	// prepare the vars needed and their memory space
	var name, address string

	person := make(map[string]string)

	/*** obtain the user input ***/
	fmt.Println("Enter a name: ")

	fmt.Scan(&name) // returns the number of items successfully scanned...

	fmt.Println("Enter an address: ")

	fmt.Scan(&address) // ...but we odn't need those return values.
	/*** user input obtained ***/

	// place the user-inputted data into a map.
	person[name] = address

	// encode it into json
	barr, _ := json.Marshal(person)

	// inform the user of what the data in Go
	fmt.Printf("\nYour data in Go's data type (a map): \n%v\n", person)

	// print the JSON for the user
	fmt.Printf("\nYour data in JSON: \n%s\n", barr)
}
