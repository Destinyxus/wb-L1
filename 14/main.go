package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable interface{}
	variable = 42

	// Determine the type of the variable
	switch v := variable.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	case bool:
		fmt.Println("Type: bool")
	case chan int:
		fmt.Println("Type: channel of int")
	default:
		fmt.Println("Unknown type:", reflect.TypeOf(v))
	}
}
