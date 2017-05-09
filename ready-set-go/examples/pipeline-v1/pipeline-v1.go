// Pipeline v1 simply has each worker function feed that next worker function via
// unbuffered channels. This is a cool pattern to use in Go but it is not doing any work
// in parallel.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Golang also features an init function besides main that is run after imports
// and global vars resolve
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Setting up all of the channels used to pass data around and inputs
	addCh := make(chan int)
	subCh := make(chan int)
	mulCh := make(chan int)
	resultCh := make(chan int)
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Kicking off the workers
	go addWorkerV1(addCh, subCh)
	go subtractWorkerV1(subCh, mulCh)
	go multiplyWorkerV1(mulCh, resultCh)

	//Feeding the initial input
	// This needs to be done in a goroutine. That is because the first couple values will get
	// pushed through the pipline just fine. But after the fourth we will have a problem.
	// Nothing will yet be pulling this of the result channel so all of the running goroutines,
	// including main will be asleep. Deadlock! If you don't believe me try taking this code
	// out of the anonymous function and seeing for yourself.
	go func() {
		for _, value := range inputs {
			addCh <- value
		}
	}()

	// Printing out the final results
	for i := 1; i <= 10; i++ {
		result := <-resultCh
		fmt.Println("=============================")
		fmt.Printf("Result %d: %d\n", i, result)
		fmt.Println("=============================")
	}

	//Cleaning up worker goroutines
	// This is not really needed for this example but I thought I would just show what should be
	// done in larger apps and what it does. All of the workers are ranging over an input channel.
	// This is an infinite loop each of those are stuck in right now. But there is a way to break
	// them out of this loop! By closing the channel we are saying that it will never be used again
	// and this tells Go that is can stop ranging over the channel and break out of the loop.
	close(addCh)
	close(subCh)
	close(mulCh)
	close(resultCh)
}

// All of the worker have a input and an output channel; they feed one another. You should pay
// special attention to the syntax of this functions signature. Not that there are <- operators on
// either side of the word chan. This provides extra type safety to the function. <-chan tell this
// this function that it can only read from this channel and chan<- means this function can only
// write to this channel. Although doing this is not required I would say that it is strongly
// encouraged unless you really mean to have a bi-directional channel.
func addWorkerV1(in <-chan int, out chan<- int) {
	for a := range in {
		// Only assigning this to a var so we can log it out. This is not necessary.
		b := a + rand.Intn(10)
		log.Printf("addWorkerV1 recieved input %d sent %d", a, b)
		out <- b
	}
}

func subtractWorkerV1(in <-chan int, out chan<- int) {
	for a := range in {
		b := a - rand.Intn(10)
		log.Printf("subtractWorkerV1 recieved input %d sent %d", a, b)
		out <- b
	}
}

func multiplyWorkerV1(in <-chan int, out chan<- int) {
	for a := range in {
		b := a * rand.Intn(10)
		log.Printf("multiplyWorkerV1 recieved input %d sent %d", a, b)
		out <- b
	}
}
