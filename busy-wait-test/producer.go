package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Producer produces data for the parent buffer
type Producer struct {
	simulator *Simulator
	PID       int // Process ID
}

// StartProducing starts the producer
func (p *Producer) StartProducing(sim *Simulator, id int) {
	p.PID = id
	p.simulator = sim

	time.Sleep(time.Duration(3 * time.Second))
	fmt.Println("Producer Starting...")
	for {
		lostCylces := 0
		for l := 0; l < p.simulator.N-1; l++ {
			p.simulator.level[p.PID] = l
			p.simulator.lastToEnter[l] = p.PID
			for p.simulator.shouldKeepWaiting(l, p.PID) || p.simulator.buffer.Size() == p.simulator.buffer.MaxSize {
				// Wait
				lostCylces++
			}
		}

		fmt.Println("Lost cycles during busy wait: ", lostCylces)
		// Critical Section Start
		fmt.Println("@@@@@@@@@@@@@@@@@ START critical section")

		value := p.simulator.nextToProduce
		fmt.Println("Inserting value ", value)
		p.simulator.buffer.Insert(value)
		p.simulator.nextToProduce++

		fmt.Println("################# END critical section")
		// Critical Section End
		p.simulator.level[p.PID] = -1

		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
	}
}
