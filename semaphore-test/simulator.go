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

	mutex *semaphore.Weighted // Mutex for concurrent access to the shared buffer
	full  *semaphore.Weighted // Full to prevent inserting into full buffer
	empty *semaphore.Weighted // Empty to prevent extracting from an empty buffer

	context context.Context //Requiered by the semaphores

	consumers []*Consumer // Consumers list
	producers []*Producer // Producers list

	nextToConsume int
	nextToProduce int
}

// Init starts the simulator with the given values
func (s *Simulator) Init(bufferSize int, producerCount int, consumerCount int, testDuration int) {
	fmt.Println("Starting Simulator...")
	// Start buffer
	s.buffer = &Buffer{
		MaxSize: bufferSize,
		data:    make([]int, bufferSize),
	}
	// Init the context variable
	s.context = context.Background()

	// Init last number
	s.nextToConsume = 0
	s.nextToProduce = 0

	// Start semaphores
	s.mutex = semaphore.NewWeighted(int64(1))
	s.full = semaphore.NewWeighted(int64(bufferSize))
	s.empty = semaphore.NewWeighted(int64(bufferSize))

	s.empty.Acquire(s.context, int64(bufferSize))

	s.consumers = []*Consumer{}
	s.producers = []*Producer{}

	// Start consumers
	for i := 0; i < consumerCount; i++ {
		fmt.Println("Adding Consumer Number ", i)
		newConsumer := &Consumer{}
		s.consumers = append(s.consumers, newConsumer)
		go newConsumer.StartConsuming(s)
	}

	// Start producers
	for i := 0; i < producerCount; i++ {
		fmt.Println("Adding Producer Number ", i)
		newProducer := &Producer{}
		s.producers = append(s.producers, newProducer)
		go newProducer.StartProducing(s)
	}

	time.Sleep(time.Duration(testDuration) * time.Second)
	fmt.Println("-----------------------------------------")
	fmt.Println("Test finished, Total inserts: ", s.nextToProduce-1, " Total Consumed: ", s.nextToConsume-1)

}
