package animal

import "fmt"

type Animal struct {
	Weight int64
}

func (a Animal) Eat() {
	fmt.Println("Animal is eating")
}

func (a Animal) Sleep() {
	fmt.Println("Animal is sleeping")
}

type Parrot struct {
	Animal
	Color string
}

func CreateParrot(color string, weight int64) Parrot {
	return Parrot{
		Animal{Weight: weight},
		color,
	}
}

func (p Parrot) Fly() {
	fmt.Println("Parrot is flying")
}

func (p Parrot) Eat() {
	fmt.Println("Parrot is eating")
}

func Main() {
	p := CreateParrot("Green", 12)
	p.Fly()
	//Parrot is flying
	p.Eat()
	//Parrot is eating
	p.Sleep()
	//Animal is sleeping
	fmt.Printf("Parrot info-> color:%s, weight:%d", p.Color, p.Weight)
	//Parrot info-> color:Green, weight:12
}
