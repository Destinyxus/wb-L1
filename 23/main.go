package main

import "fmt"

func main() {
	fmt.Println(deleteElement([]int{1, 2, 3, 4, 5, 6}, 2))
}

func deleteElement(arr []int, index int) []int {

	for i := index + 1; i < len(arr); i++ {
		arr[i-1] = arr[i]

	}

	return arr[:len(arr)-1]
}
