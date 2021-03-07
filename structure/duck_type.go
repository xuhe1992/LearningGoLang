package structure

import "fmt"

type DuckType interface {
	walk()
	say()
}

type A struct {
	Name string
	Age int
}

type B struct {
	name string
	age int
}

func (a *A) walk() {
	fmt.Println("A walking...")
}

func (a *B) say() {
	fmt.Println("B saying...")
}

func TestDuck(d DuckType) {
	d.walk()
	d.say()
}
