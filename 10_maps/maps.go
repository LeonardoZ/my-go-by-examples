package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("emp: ", m)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1 ", v1)
	fmt.Println("len ", len(m))

	delete(m, "k2")
	fmt.Println("after delete  ", m)

	_, psr := m["k2"]
	fmt.Println("prs: ", psr)

	n := map[string]int{"foo": 1, "bar": 1}
	fmt.Println("map: ", n)
}
