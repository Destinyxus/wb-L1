package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{43, 21, 65, 22, 64, 77, 32}))
}

func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivotElement := arr[pivotIndex]

	left := []int{}
	right := []int{}

	for i := 0; i < len(arr); i++ {
		if i == pivotIndex {
			continue
		}
		if pivotElement > arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quickSort(left), pivotElement), quickSort(right)...)

}
