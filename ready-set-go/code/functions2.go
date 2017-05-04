package main

import "fmt"

func main() {
	fmt.Println(multipleReturns())
}

// START OMIT
func multipleReturns() (int, bool) {
	return 2, true
}

// END OMIT
