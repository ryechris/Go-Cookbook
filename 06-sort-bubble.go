package main

import "fmt"

func main() {

	// remember, have to specify type. Precise.
	// declare variables to hold numbers from user.
	var a, b, c, d, e int

	// make slice to hold input from user.
	slice := make([]int, 0)
	// the fmt.Scan() function returns num, err; where numb is the number of bytes read,
	// and err is the error (if any). We don't need these returns for now.

	/** Obtain number**/
	fmt.Println("Please enter the first number: ")
	fmt.Scan(&a)
	slice = append(slice, a)

	fmt.Println("Please enter the second number: ")
	fmt.Scan(&b)
	slice = append(slice, b)

	fmt.Println("Please enter the third number: ")
	fmt.Scan(&c)
	slice = append(slice, c)

	fmt.Println("The fourth: ")
	fmt.Scan(&d)
	slice = append(slice, d)

	fmt.Println("And the 5th number: ")
	fmt.Scan(&e)
	slice = append(slice, e)
	/****/

	fmt.Printf("The numbers you just inputted,")
	fmt.Printf("in the order you entered them: %v\n\n", slice)
	fmt.Printf("Sorting...\n\n")

	last_index := len(slice) - 1

	// We set the slice's last index as the limit of how many times the loops iterate.
	// so if the last index is 1, the outer loop has only 1 iteration.
	// if the last index is 5, then the outer loop goes for 5 iterations.
	for upper_bound := last_index; upper_bound > 0; upper_bound-- { // (n-1) operations, where n = len(slice)

		for i := 1; i <= upper_bound; i++ { // (n-1) operations, then (n-2), ... (n-(n-1)=1) for the last loop.
			if slice[i] < slice[i-1] {
				swap(slice, i)
			}
		} // end of inner loop

		fmt.Printf("At the end of this iteration, the slice is %v\n", slice)
		fmt.Printf("end of iteration %v\n\n", upper_bound)
	} // end of outer loop

	fmt.Printf("sorted slice: %v\n", slice)
	fmt.Println("note: if letter was inputted, it was replaced with 0.")
}

func swap(slice []int, index int) {
	// swap(i, 0) fails => has to start at index 1.
	slice[index-1], slice[index] = slice[index], slice[index-1]
}
