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

func (d Dog) Name() string { // Value receiver
	return d.name
}
func (d Dog) FavoriteToy() string { //Pointer receiver would start with (d *Dog)
	return d.favoriteToy
}

func PetAnimal(p Pet) { // HL
	fmt.Printf("You petted %s!\n", p.Name())
}

func main() {
	d := Dog{"Meeko", "🍖"}
	PetAnimal(d)
}

// END OMIT
