package main

// START OMIT
import (
	"fmt"
	//"time"
)

func printNum(i int) {
	fmt.Println(i)
}

func printLoop() {
	for i := 0; i <= 10; i++ {
		go printNum(i) // HL
	}
}

func main() {
	go printLoop() // HL
	//time.Sleep(100 * time.Millisecond)
}

// END OMIT
