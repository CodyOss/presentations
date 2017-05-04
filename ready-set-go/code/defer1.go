package main

import "fmt"

func main() {
	doSomething()
}

// START OMIT
func doSomething() {
	defer fmt.Println("Do something like closing a resource") // HL
	fmt.Println("Working on a resource")
}

// END OMIT
