package main

import "fmt"

func main() {
	// START OMIT
	feeling := "happy"

	switch feeling {
	case "sad":
		fmt.Println("You are sad! 😭")
	case "happy":
		fmt.Println("You are happy! 😁")
		//fallthrough
	default:
		fmt.Println("I am not sure what you are... 🙃")
	}
	// END OMIT
}
