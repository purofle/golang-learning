package main

import "fmt"

type Shit struct {
	Color string
	Taste string
}

type Shit2 struct {
	Shit
}

func (s Shit) eat() string {
	return fmt.Sprintf("You ate a %s shit! It smells like %s", s.Color, s.Taste)
}

func (s *Shit) init() *Shit {
	return &Shit{Color: "green", Taste: "rotten eggs"}
}

func test(nmsl int) func(int) int {
	return func(nmsl int) int {
		return nmsl
	}
}

func main() {
	fmt.Println(Shit{Color: "green", Taste: "rotten eggs"}.eat())

	var i interface{}

	shit := new(Shit)

	fmt.Printf("%T\n", shit.Color)
	fmt.Println(shit.Color)

	i = nil
	i = 1

	fmt.Println(i)
}
