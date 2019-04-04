package main

import (
	"flag"
	"fmt"
)

func main() {
	workdPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 43, "an int")
	boolPtr := flag.Bool("fork", false, " a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *workdPtr)
	fmt.Println("number:", *numbPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}
