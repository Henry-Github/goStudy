package main

import "fmt"

// struct 结构体 methods

type Animal struct {
	name string
	legs int
}

func (*Animal) voice() {
	fmt.Println("~~~~~~~~")
}

func (animal *Animal) changeNamePtr(name string) {
	animal.name = name
}

func (animal Animal) changeName(name string) {
	// 实际值并不会被改变
	animal.name = name
}

func initAnimal(name string, legs int) *Animal {
	animal := &Animal{
		name,
		legs,
	}
	return animal
}

func main() {
	cat := initAnimal("cat", 4)
	cat.voice()
	fmt.Println(cat.name)
	cat.changeNamePtr("pig")
	fmt.Println(cat.name)
	cat.changeName("monkey")
	fmt.Println(cat.name)

}
