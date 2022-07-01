package main

import "fmt"

type Animal int

const (
	Turtle Animal = iota
	Pig
	Horse
	Cheetah
)

func (a Animal) String() string {
	return [...]string{"Turtle", "Pig", "Horse", "Cheetah"}[a]
}

func main() {
	var animal Animal = Cheetah
	ShowAnimalInfo(animal)
	_a := Pig
	ShowAnimalInfo(_a)
	_a = Horse
	ShowAnimalInfo(_a)

}
func ShowAnimalInfo(a Animal) {
	// that function show results info about animals
	fmt.Print("That Animal is ", a)
	switch a {
	case 0:
		fmt.Println(" The Slowest.")
	case 1:
		fmt.Println(" Slow.")
	case 2:
		fmt.Println(" Fast.")
	case 3:
		fmt.Println(" The Fastest.")
	default:
		fmt.Println(" stays.")
	}
}
