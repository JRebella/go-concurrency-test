package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// Simulator a
type Simulator struct {
	buffer *Buffer

	level       []int
	lastToEnter []int

	N int // Number of consumers + producers

	mutex *semaphore.Weighted // Mutex for concurrent access to the level arrays

	context context.Context //Requiered by the semaphore

	consumers []*Consumer // Consumers list
	producers []*Producer // Producers list

	nextToConsume int
	nextToProduce int
}

// Init starts the simulator with the given values
func (s *Simulator) Init(bufferSize int, producerCount int, consumerCount int) {
	fmt.Println("Starting Simulator...")
	// Start buffer
	s.buffer = &Buffer{
		MaxSize: bufferSize,
		data:    make([]int, bufferSize),
	}
	// Init the control arrays
	s.N = producerCount + consumerCount
	s.level = make([]int, s.N)
	for i := range s.level {
		s.level[i] = -1
	}
	s.lastToEnter = make([]int, s.N-1)
	for i := range s.lastToEnter {
		s.lastToEnter[i] = -1
	}

	// Init the context variable
	s.context = context.Background()

	// Init last number
	s.nextToConsume = 0
	s.nextToProduce = 0

	// Start semaphores
	s.mutex = semaphore.NewWeighted(int64(1))

	s.consumers = []*Consumer{}
	s.producers = []*Producer{}

	PID := 0

	fmt.Println("N: ", s.N)
	// Start producers
	for i := 0; i < producerCount; i++ {
		fmt.Println("Adding Producer Number ", PID)
		newProducer := &Producer{}
		s.producers = append(s.producers, newProducer)
		go newProducer.StartProducing(s, PID)
		PID++
	}

	// Start consumers
	for i := 0; i < consumerCount; i++ {
		fmt.Println("Adding Consumer Number ", PID)
		newConsumer := &Consumer{}
		s.consumers = append(s.consumers, newConsumer)
		go newConsumer.StartConsuming(s, PID)
		PID++
	}

	time.Sleep(10 * time.Second)
	fmt.Println("-----------------------------------------")
	fmt.Println("Test finished, Total inserts: ", s.nextToProduce, " Total Consumed: ", s.nextToConsume)
}

func (s *Simulator) shouldKeepWaiting(level int, PID int) bool {
	if s.lastToEnter[level] != PID {
		return false
	}
	for k := range s.level {
		if s.level[k] >= level {
			if k != PID {
				return true
			}
		}
	}
	return false
}
