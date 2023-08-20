// Copyright (c) 2023 Riyan Christy. All rights reserved.
// This program computes factorial.
// Factorial is a relatively compute-intensive calculation.
// We will use a total of 7 threads (goroutines; including the Main one)
// to compute the factorial of a number.
// We will use this factorial example to illustrate
// goroutines, channels and so forth.

package main

import (
	"fmt"
)

func main() {
	var ui int
	fmt.Println("> Hi! Welcome to the Factorial Calculator")
	fmt.Println("> Please enter the number you want the factorial of:")
	fmt.Println("> e.g. enter 5 for 5! (but < 21; so 20 max.)")
	fmt.Printf("> ")
	_, err := fmt.Scan(&ui) // errors as values. 

	// Handle errors
	if err != nil {
		fmt.Println("Please enter a natural number (1,2,3,4,5...)")
	} else if ui < 0 {
		fmt.Println("Please enter a number greater than 0.")
		return
	} else if ui == 0 {
		fmt.Println("1")
		return
	}

	// n! = n * (n-1) * (n-2) ... * 1
	// So we want to list those values here in an array:
	slice := make([]int, 0)
	for i := 1; i <= ui; i++ {
		slice = append(slice, i)
	}

	// We're going to make a channel to minimize blocking in sends/receives.
	c := make(chan int, 4)

	/* divide slice into four, each gets a goroutine */
	u := len(slice) / 4

	i, slice := slice[:u], slice[u:]
	ii, slice := slice[:u], slice[u:]
	iii, iv := slice[:u], slice[u:]
	slice = slice[len(slice):] // empty the slice for further use.
	slices := [][]int{i, ii, iii, iv}
	/* slice division concludes */

	routineLoop(slices, c)

	// We put the <-c in its own loop, because in case it blocks,
	// it doesn't block the calling of the next go statement.
	for j := 0; j < len(slices); j++ {
		slice = append(slice, <-c)
	}

	// receive blocks if the buffer is empty;
	// so by this time, buffer has been emptied and slice has our values.
	u = len(slice) / 2
	i, ii = slice[:u], slice[u:]
	slice = slice[len(slice):]
	slices = [][]int{i, ii}

	routineLoop(slices, c)

	a := <-c
	b := <-c

	fmt.Printf("> The factorial: %v! = %v\n", ui, a*b)

	// side note:
	// There are a lot of variables that we could "delete" after use,
	// But we can let Go's garbage collection handle that.
}

func factorial(s []int, c chan int) {
	f := 1
	for _, v := range s {
		f = v * f
	}
	c <- f
}

func routineLoop(slices [][]int, ch chan int) {
	for index := 0; index < len(slices); index++ {
		go factorial(slices[index], ch)
	}
}
