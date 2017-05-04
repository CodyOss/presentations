package main

import "fmt"

func main() {
	// START OMIT
	msgs := make(chan string, 2)

	msgs <- "hello"
	msgs <- "you"
	//msgs <- "there"

	fmt.Println(<-msgs)
	//END OMIT
}
