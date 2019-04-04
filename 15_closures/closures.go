package main

import (
	"fmt"
	"strconv"
)

func intSeq() func() int {
	state := "zero"
	i := 0
	return func() int {
		i++
		state = strconv.Itoa(i)
		fmt.Println("conv: ", state)
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}
