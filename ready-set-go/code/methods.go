package main

import "fmt"

type Dog struct {
	Name        string
	Breed       string
	age         int
	favoriteToy string
}

// START OMIT
func (d Dog) DogYears() int { //Value receiver
	return d.age * 7
}

func (d *Dog) FavoriteToy() string { //Pointer receiver
	return d.favoriteToy
}

func main() {
	d := Dog{
		Name:        "Meeko",
		Breed:       "Great Pyrenees",
		age:         4,
		favoriteToy: "🍖",
	}

	fmt.Printf("%s's is %d years old in dog Years!\n", d.Name, d.DogYears())
	dp := &d
	fmt.Printf("And his favorite toy is %s!", dp.FavoriteToy())
}

// END OMIT
