package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Consumer consumes data from the parent buffer
type Consumer struct {
	simulator *Simulator
	PID       int // Process ID
}

// StartConsuming starts the consumer
func (c *Consumer) StartConsuming(sim *Simulator, id int) {
	c.PID = id
	c.simulator = sim

	time.Sleep(time.Duration(3 * time.Second))
	fmt.Println("Consumer Starting...")
	for {
		lostCylces := 0
		for l := 0; l < c.simulator.N-1; l++ {
			c.simulator.level[c.PID] = l
			c.simulator.lastToEnter[l] = c.PID
			for c.simulator.shouldKeepWaiting(l, c.PID) || c.simulator.buffer.Size() == 0 {
				// Wait
				lostCylces++
			}
		}

		fmt.Println("Lost cycles during busy wait: ", lostCylces)
		// Critical Section Start
		fmt.Println("@@@@@@@@@@@@@@@@@ START critical section")
		extractedValue := c.simulator.buffer.Extract()

		if extractedValue != c.simulator.nextToConsume {
			panic("Panic")
		}

		fmt.Println("Consuming value ", extractedValue)
		c.simulator.nextToConsume++

		fmt.Println("################# END critical section")
		// Critical Section End
		c.simulator.level[c.PID] = -1

		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
	}
}
