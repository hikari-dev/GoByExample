package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{age: 30, name: "Alice"})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{"Ann", 40})
	fmt.Println(newPerson("Jhon"))
	s := person{name: "Sean", age: 22}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}
