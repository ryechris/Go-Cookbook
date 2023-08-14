package main

import "fmt"

// This function computes the displacement of an object with constant acceleration
// displacement = (1/2)a(t^2) + v(t) + s
func computeDisplacement(a, v, s float64) func(float64) float64 {
	// a: constant acceleration, v: initial velocity, s: initial displacement

	f := func(t float64) float64 { // f(t) function of time.
		return (float64(0.5)*a*t*t + v*t + s)
	}

	return f
}

func main() {
	// Declare vars & type to hold vals for the
	// constant acceleration, initial velocity, initial displacement, and time:
	var a, v, s, t float64

	fmt.Println("\nAll units are in Système International meters & seconds.\nNo need to enter units; just the number.\n")
	fmt.Printf("constant acceleration (e.g. gravity: 9.8): ")
	fmt.Scan(&a)

	fmt.Printf("initial velocity (e.g. 0): ")
	fmt.Scan(&v)

	fmt.Printf("initial displacement (e.g. 0): ")
	fmt.Scan(&s)

	fmt.Println("\nOK, got it. So the values are: ")
	fmt.Printf("constant acceleration: %v\n", a)
	fmt.Printf("initial velocity: %v\n", v)
	fmt.Printf("initial displacement: %v\n", s)
	fmt.Println("\n")

	f := computeDisplacement(a, v, s)

	fmt.Println("Now, you want to calculate the displacement after how many seconds?")
	fmt.Println("e.g. 10.5 for diplacement after 10 and a half seconds.")
	fmt.Printf("So, after how many seconds: ")
	fmt.Scan(&t)

	fmt.Printf("The displacement is f(t) = f(%v) = %v meters\n", t, f(t))
}

