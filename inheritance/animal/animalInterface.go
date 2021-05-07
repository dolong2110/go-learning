package animal

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Bird struct {
}

type CircusBird struct {
}

type flyer struct {
	Animal Animal
}

func (b Bird) Eat() {
	fmt.Println("Bird is eating")
}

func (b Bird) Sleep() {
	fmt.Println("Bird is sleeping")
}

func (c CircusBird) Eat() {
	fmt.Println("CircusBird is eating")
}

func Main() {
	/* Some code */
	p := flyer{}
	p.Animal = Bird{}
	p.Animal.Eat()
	// Bird is eating
	/* Maybe our Parrot want to show some stunt now */
	/* CircusBird is just another struct on Animal Interface */
	p.Animal = CircusBird{}
	p.Animal.Eat()
	// CircusBird is eating
}
