package main

import "fmt"

func main() {
	fmt.Println("\nStarting tests...")

	sim := &Simulator{}
	sim.Init(50000, 255200, 300003)
}
