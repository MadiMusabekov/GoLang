package main

import "fmt"

type Socket interface {
	Charge(s string)
}

type Smartphone struct {
	Name  string
	Price int
}

func (s Smartphone) GetName() string {
	return s.Name
}

func (s Smartphone) Charge(ss string) {
	fmt.Println("Charging ", s)
}

func main() {
	s := new(Smartphone)
	s.Name = "iPhone"
	s.Price = 450
	fmt.Println(s)
}
