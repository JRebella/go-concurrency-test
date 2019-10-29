package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Producer produces data for the parent buffer
type Producer struct {
	simulator *Simulator
}

// StartProducing starts the producer
func (p *Producer) StartProducing(sim *Simulator) {
	time.Sleep(time.Duration(3 * time.Second))

	fmt.Println("Producer Starting...")
	p.simulator = sim
	for {
		p.simulator.full.Acquire(p.simulator.context, 1)  //Block if buffer is full
		p.simulator.mutex.Acquire(p.simulator.context, 1) //Block if buffer is being modified

		value := p.simulator.nextToProduce
		fmt.Println("<<<<<Inserting value ", value)
		p.simulator.buffer.Insert(value)
		p.simulator.nextToProduce++

		p.simulator.mutex.Release(1) //Signal buffer is free to use
		p.simulator.empty.Release(1) //Signal buffer has data to be consumed

		time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)

	}
}
