package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("1000000000000000000000", 10)
	b := new(big.Int)
	b.SetString("200000000000000000000", 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Int).Quo(a, b)
	
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", diff.String())
	fmt.Printf("a * b = %s\n", prod.String())
	fmt.Printf("a / b = %s\n", quot.String())
}