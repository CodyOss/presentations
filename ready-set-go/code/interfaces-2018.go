package main

import "fmt"

// START OMIT
type Pet interface { // HL
	Name() string        // HL
	FavoriteToy() string // HL
}

type Dog struct { // Implicitly a Pet because it fulfills the interface
	name        string
	favoriteToy string
}

func (d Dog) Name() string {
	return d.name
}
func (d Dog) FavoriteToy() string {
	return d.favoriteToy
}

func PetAnimal(p Pet) { // HL
	fmt.Printf("You petted %s!\n", p.Name())
}

func main() {
	d := Dog{"Sven", "🍖"}
	PetAnimal(d)
}
// END OMIT
