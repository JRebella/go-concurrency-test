package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Consumer consumes data from the parent buffer
type Consumer struct {
	simulator *Simulator
}

// StartConsuming starts the consumer
func (c *Consumer) StartConsuming(sim *Simulator) {
	time.Sleep(time.Duration(3 * time.Second))

	fmt.Println("Consumer Starting...")
	c.simulator = sim
	for {
		c.simulator.empty.Acquire(c.simulator.context, 1) //Block if buffer is empty
		c.simulator.mutex.Acquire(c.simulator.context, 1) //Block if buffer is being modified

		extractedValue := c.simulator.buffer.Extract()

		if extractedValue != c.simulator.nextToConsume {
			panic("Panic")
		}

		fmt.Println(">>>>>Consuming value ", extractedValue)
		c.simulator.nextToConsume++

		c.simulator.mutex.Release(1) //Signal buffer is free to use
		c.simulator.full.Release(1)  //Signal buffer has room for more data

		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
	}
}
