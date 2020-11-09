package main

import "fmt"

// Prints everything
type Printable interface {
	print()
}

type person struct {
	name string
}

type figure struct {
	name string
}

func (f *figure) print() {
	fmt.Println("[figure]", f.name)
}

// Methods of a class, receiver arguments. (p person) makes a copy of the instance
func (p person) print() {
	fmt.Println("[Person]", p.name)
}

// Methods of a class, receiver arguments. (p *person) copy of pointer of p
func (p *person) clean() {
	p.name = ""
}

func invokePrint(p Printable) {
	p.print()
}

func main() {
	p := &person{name: "Juan"}
	invokePrint(p)

	f := &figure{"Figure"}
	invokePrint(f)
}
