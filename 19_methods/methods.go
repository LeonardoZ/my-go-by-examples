package main

import "fmt"

type rect struct {
	width, height int
}

// pointer value
func (r rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

// pointer receiver
func (r *rect) set(height, width int) {
	r.width = width
	r.height = height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area ", r.area())
	fmt.Println("perim ", r.perim())

	rp := &r
	fmt.Println(rp)
	fmt.Println("area ", rp.area())
	fmt.Println("perim ", rp.perim())

	r.set(10, 10)

	fmt.Println("area ", r.area())
	fmt.Println("perim ", r.perim())

}
