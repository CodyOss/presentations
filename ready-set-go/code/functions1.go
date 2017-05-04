package main

import "fmt"

func main() {
	fmt.Println(multiply(2, 3))
}

// START OMIT
// input will be 1 and 2
func multiply(a, b int) int {
	return a * b
}

// END OMIT
