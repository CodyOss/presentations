package main

import (
	"fmt"
	"time"
)

// START OMIT
func pullWorkOffChannels(ch1, ch2 chan string) {
	for { // A while loop in Go (aka unbound for loop)
		select { // HL
		case msg := <-ch1: // HL
			fmt.Println("CH1: " + msg)
		case msg := <-ch2: // HL
			fmt.Println("CH2: " + msg)
		case <-time.After(3 * time.Second): // HL
			return
		}
	}
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { time.Sleep(time.Second * 1); ch1 <- "🔥" }()
	go func() { time.Sleep(time.Second * 2); ch2 <- "🌊" }()
	pullWorkOffChannels(ch1, ch2)
	fmt.Println("Done")
	// END OMIT
}
