package main

import "fmt"

func main() {
	// START OMIT
	msg := make(chan string) // HL

	go func() { msg <- "hello" }()

	receive := <-msg // HL
	fmt.Println(receive)
	//END OMIT
}
