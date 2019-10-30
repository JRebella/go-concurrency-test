package main

import "fmt"

func main() {
	fmt.Println("\nStarting tests...")

	sim := &Simulator{}
	// Recommended: 3-2-2
	sim.Init(3, 2, 2, 40)
}
