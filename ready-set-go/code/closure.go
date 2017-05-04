package main

import "fmt"

// START OMIT
func next() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := next()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

// END OMIT
