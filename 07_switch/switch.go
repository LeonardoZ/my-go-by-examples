package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("Write ", i, " as")

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a Weekend")
	default:
		fmt.Println("It's a Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("A bool")
		case int:
			fmt.Println("A int")
		default:
			fmt.Println("Have no idea ", t)
		}

	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")

}
