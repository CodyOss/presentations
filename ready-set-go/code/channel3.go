package main

import (
	"fmt"
	"time"
)

// START OMIT
func doSomeWork(done chan bool) {
	fmt.Println("Doing some work")
	time.Sleep(1 * time.Second)
	fmt.Println("And I am done")
	close(done) // HL
}

func main() {
	done := make(chan bool)

	go doSomeWork(done)
	<-done // HL
}

//END OMIT
