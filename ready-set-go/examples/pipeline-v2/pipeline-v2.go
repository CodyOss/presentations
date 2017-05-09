// Pipeline v2 take everything learned in v1 and shows how to tweak that pattern to make things run
// run in parallel. This is done be creaking groups of works and buffered channels.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var size int

func init() {
	rand.Seed(time.Now().UnixNano())
	// The flag package in Go is great for build CLIs or apps that want command line options
	flag.IntVar(&size, "size", 10, "how large to make buffered channels and worker pools")
	flag.Parse()
}

// Making a type so the signature to buildWorkerGroup does not look as bad
type workerFn func(<-chan int, chan<- int, *sync.WaitGroup)

func main() {
	// Setting up all of the channels used to pass data around and inputs
	addCh := make(chan int, size)
	subCh := make(chan int, size)
	mulCh := make(chan int, size)
	resultCh := make(chan int)

	// Waitgroups are used to keep track of goroutines. We will later call Wait on them latter.
	// This is a blocking action and will wait until Done has been called to each item added to the
	// WaitGroup.
	addWg := new(sync.WaitGroup)
	subWg := new(sync.WaitGroup)
	mulWg := new(sync.WaitGroup)
	inputs := make([]int, 1000)
	for i := 1; i <= 1000; i++ {
		inputs = append(inputs, i)
	}

	// Kicking off the worker groups
	// This could be done in a goroutine too if you really wanted.
	buildWorkerPool(addWorkerV2, addCh, subCh, addWg)
	buildWorkerPool(subtractWorkerV2, subCh, mulCh, subWg)
	buildWorkerPool(multiplyWorkerV2, mulCh, resultCh, mulWg)

	//Feeding the initial input
	go func() {
		for _, value := range inputs {
			addCh <- value
		}
		// All inputs have been placed on the channel already so we know it is now safe to close this channel.
		close(addCh)
	}()

	// Cleaning up worker goroutines
	// Before closing channels we need to make sure that all goroutines are done working. Closing
	// a buffered channel will allow it to drain remaining items already in the channel if
	// something is reading from it.
	go func() {
		addWg.Wait()
		close(subCh)
		subWg.Wait()
		close(mulCh)
		mulWg.Wait()
		close(resultCh)
	}()

	// Printing out the final results
	for i := 1; i <= 1000; i++ {
		result := <-resultCh
		fmt.Println("=============================")
		fmt.Printf("Result %d: %d\n", i, result)
		fmt.Println("=============================")
	}
}

// Notice how the worker funcs are implicitly a workerFn that we declared at the top. This is
// just a helper method to build the pools of channels
func buildWorkerPool(worker workerFn, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < size; i++ {
		go worker(in, out, wg)
		wg.Add(1)
	}
}

func addWorkerV2(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for a := range in {
		b := a + rand.Intn(10)
		log.Printf("addWorkerV1 recieved input %d sent %d", a, b)
		out <- b
	}
	wg.Done()
}

func subtractWorkerV2(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for a := range in {
		b := a - rand.Intn(10)
		log.Printf("subtractWorkerV1 recieved input %d sent %d", a, b)
		out <- b
	}
	wg.Done()
}

func multiplyWorkerV2(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for a := range in {
		b := a * rand.Intn(10)
		log.Printf("multiplyWorkerV1 recieved input %d sent %d", a, b)
		out <- b
	}
	wg.Done()
}
