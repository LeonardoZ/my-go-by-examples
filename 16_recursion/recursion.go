package main

import "fmt"
import "math/big"

func fact(n *big.Int) *big.Int {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	if n.Cmp(zero) == 0 {
		return one
	}
	var minusOne = new(big.Int).Sub(n, one)
	return n.Mul(n, fact(minusOne))
}

func main() {
	var factorialOf *big.Int = big.NewInt(30)
	res := fact(factorialOf)
	fmt.Println("Factorial is \n", res)
}
