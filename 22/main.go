package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	a := big.NewInt(int64(math.Pow(2, 20)))
	b := big.NewInt(int64(math.Pow(2, 20)))

	b2 := &big.Int{}
	fmt.Println(b2.Add(a, b))
	fmt.Println(b2.Div(a, b))
	fmt.Println(b2.Mul(a, b))
	fmt.Println(b2.Sub(a, b))

}
