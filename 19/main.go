package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse("рыба"))
}

func reverse(word string) string {
	a := []rune(word)
	arr := make([]rune, len(a))
	first := 0

	for i := len(a) - 1; i >= 0; i-- {
		arr[first] = a[i]
		first++
	}

	return string(arr)
}
