// Pipeline v3 is more like v1 than v2. The difference here is that instread of having the workers
// receive the data directly from the previous worker, the main goroutine will send the data to
// the next worker with a select statment. This is a contrived example to use select, you might
// really want to do this if the other goroutine has access to something that our workers don't.
// In this example we are just using it so all the logging happens in one place.
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
	addInCh := make(chan int)
	addOutCh := make(chan int)
	subInCh := make(chan int)
	subOutCh := make(chan int)
	mulInCh := make(chan int)
	mulOutCh := make(chan int)
	resultCh := make(chan int)
	done := make(chan bool)
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Kicking off the workers
	go addWorkerV3(addInCh, addOutCh)
	go subtractWorkerV3(subInCh, subOutCh)
	go multiplyWorkerV3(mulInCh, mulOutCh)

	//Feeding the initial input
	go func() {
		for _, value := range inputs {
			addInCh <- value
		}
	}()
	// Routing data
	// The below commented out code leads to a deadlock. A select is needed for pair of ins and
	// outs. Selects are meants to read from multiple goroutines but not if each of those cases
	// is doing routing logic to different channels
	/*go func() {
		for {
			select {
			case val := <-addOutCh:
				log.Printf("addWorker sent back: %d", val)
				subInCh <- val
			case val := <-subOutCh:
				log.Printf("subWorker sent back: %d", val)
				mulInCh <- val
			case val := <-mulOutCh:
				log.Printf("mulWorker sent back: %d", val)
				resultCh <- val
			case <-time.After(1 * time.Millisecond):
			case <-done:
				return
			}
		}
	}()*/
	go func() {
		for {
			select {
			case val := <-addOutCh:
				log.Printf("addWorker sent back: %d", val)
				subInCh <- val
			case <-done:
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case val := <-subOutCh:
				log.Printf("subWorker sent back: %d", val)
				mulInCh <- val
			case <-done:
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case val := <-mulOutCh:
				log.Printf("mulWorker sent back: %d", val)
				resultCh <- val
			case <-done:
				return
			}
		}
	}()

	// Printing out the final results
	for i := 1; i <= 10; i++ {
		result := <-resultCh
		fmt.Println("=============================")
		fmt.Printf("Result %d: %d\n", i, result)
		fmt.Println("=============================")
	}
	// This is to break out of the select to insure there is no data race and freeing up that
	// goroutine.
	close(done)

	//Cleaning up worker goroutines
	closeAll(resultCh, addInCh, addOutCh, subInCh, subOutCh, mulInCh, mulOutCh)
}

func addWorkerV3(in <-chan int, out chan<- int) {
	for a := range in {
		out <- a + rand.Intn(10)
	}
}

func subtractWorkerV3(in <-chan int, out chan<- int) {
	for a := range in {
		out <- a - rand.Intn(10)
	}
}

func multiplyWorkerV3(in <-chan int, out chan<- int) {
	for a := range in {
		out <- a * rand.Intn(10)
	}
}

// Go supports variadic funcs
func closeAll(chans ...chan int) {
	for _, ch := range chans {
		close(ch)
	}
}
