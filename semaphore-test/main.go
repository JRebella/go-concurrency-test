package main

import "fmt"

func main() {
	fmt.Println("\nStarting tests...")

	sim := &Simulator{}
	sim.Init(3, 2, 2)
}
