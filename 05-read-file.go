/* riyan [at] linux.com */

package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	// make the slice to hold the structs
	slice := make([]Person, 0)

	// file content will be placed into this array:
	array_of_bytes := make([]byte, 20)

	// make var to store input from user
	var input int

	for {
		// Prompt the user to give the file name to open (example1.txt or example2.txt)
		fmt.Println("Please choose the file you want open: ")
		fmt.Println("(Enter 1 for file1, or 2 for file2)")
		fmt.Printf("So, 1 or 2: ")
		_, e := fmt.Scan(&input)
		if e != nil {
			fmt.Println("Please try again with a number (1 or 2).")
		} else if input < 1 || input > 2 {
			fmt.Println("We have only two files: 1 or 2.")
		} else { // here the user has entered 1 or 2. So we have to deliver the resuult.
			switch input {
			case 1:
				// open file
				file, _ := os.Open("05-file1.txt")

				// read file
				slice = read(array_of_bytes, file, slice)

				// close file
				file.Close()
			case 2:
				// open file
				fmt.Println("Yassir")
				file, _ := os.Open("05-file2.txt")

				slice = read(array_of_bytes, file, slice)

				// close file
				file.Close()
			}
			slice = slice[:len(slice)-1]

			// print results to user
			// by iterating through resultant slice
			for _, value := range slice {
				fmt.Printf("First Name: %s\n", value.fname)
				fmt.Printf("Last Name: %s \n\n", value.lname)
			}
		}
	}
}

func read(array_of_bytes []byte, file *os.File, slice []Person) []Person {
	for _, err := file.Read(array_of_bytes); err == nil; _, err = file.Read(array_of_bytes) {

		// we could split this into several vars, but we only need the array.
		two_names := strings.Split(strings.TrimSpace(string(array_of_bytes)), ",")

		// append the slice
		slice = append(slice, Person{fname: two_names[0], lname: two_names[1]})
	}
	return slice // return slice lest the main slice above doesn't get updated.
}

