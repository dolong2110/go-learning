package animal

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Bird struct {
}

func (b Bird) Eat() {
	fmt.Println("Bird is eating")
}

func (b Bird) Sleep() {
	fmt.Println("Bird is sleeping")
}

func Main(){
	/* Some code */
	p.Animal = Bird{}
	p.Eat()
	// Bird is eating
	/* Maybe our Parrot want to show some stunt now */
	/* CircusBird is just another struct on Animal Interface */
	p.Animal = CircusBird{}
	p.Eat()
	// CircusBird is eating
}