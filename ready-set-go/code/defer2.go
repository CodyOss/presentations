package main

import "fmt"

func main() {
	doSomethingInALoop()
}

// START OMIT
func doSomethingInALoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // HL
	}
}

// END OMIT
