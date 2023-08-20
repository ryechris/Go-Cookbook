// Copyright (c) 2023 Riyan Christy. All rights reserved.
// This program is an implementation of Dijkstra's Resource Hierarchy solution,
// (After all, the Dining Philosophers was originally an exam problem
// he created for his students.)
// Since this is Go, Riyan has modified a few things a little,
// but the essence is the same (as you can examine yourself).
package main

import (
	"fmt"
	"sync"
)

// making waitgroup global
// refactor later
var wg sync.WaitGroup

type Chopstick struct {
	mu sync.Mutex
}

type Philosopher struct {
	FirstChopstick, SecondChopstick *Chopstick // pointer type to chopstick
	number                          int
}

func (p *Philosopher) eat() {
	// The infinite for loop has been commented out,
	// so that each philosopher eats once and then exits.
	// However, this solution has no problems -- even for an infinite loop.
	// If you want to test it, just remove the
	// comments the for loop brackets (on lines 33 and 43).
	// for {
	p.FirstChopstick.mu.Lock()
	p.SecondChopstick.mu.Lock()

	fmt.Printf("Philosopher %v eating\n", p.number)

	p.FirstChopstick.mu.Unlock()
	p.SecondChopstick.mu.Unlock()

	fmt.Printf("Philosopher %v exiting...\n\n", p.number)
	// }
	wg.Done()
}

func main() {
	// slice to hold the chopsticks
	Chopsticks := make([]*Chopstick, 5)

	// fill the slice with new chopsticks.
	for i := 0; i < 5; i++ {
		Chopsticks[i] = new(Chopstick)
	}

	// slice for the five philosophers
	philosophers := make([]*Philosopher, 5)

	// bring in the philosophers
	for i := 0; i < 5; i++ {
		var f, s *Chopstick // f for first, s for second

		// Here we set the Resource Hierarchy: we number each chopstick.
		// (use %5 because without it, p4 would have chopsticks[5].)
		if ((i + 1) % 5) < i {
			f = Chopsticks[(i+1)%5]
			s = Chopsticks[i]
		} else {
			f = Chopsticks[i]
			s = Chopsticks[(i+1)%5]
		}

		// we initialize a philosopher with a number, so we can differentiate.
		philosophers[i] = &Philosopher{f, s, i + 1}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ { // 5 goroutines here
		go philosophers[i].eat()
	}
	wg.Wait()

	fmt.Println("SOLUTION SUCCEEDS")
	fmt.Println("-- 6 goroutines (including Main's), 0 deadlock.")
}

