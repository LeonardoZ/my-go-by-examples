package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bog", 20})
	fmt.Println(person{"Alice", 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{"Ann", 50})

	s := person{name: "Sean", age: 45}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}
