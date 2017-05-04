package main

import "fmt"

// START OMIT
func saveMe() {
	if r := recover(); r != nil { // HL
		fmt.Println("I saved you from your silly panicing!")
	}
}

func panicNow() {
	fmt.Println("AHHH, I am Panicing!")
	panic("HAHAHA") // HL
}

func doSomeWork() {
	//defer saveMe()
	panicNow()
}

func main() {
	doSomeWork()
	fmt.Println("Things are good again 😊")
}

// END OMIT
