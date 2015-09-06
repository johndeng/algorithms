package main

import (
	"fmt"
	"math/rand"
)

func quickSort(items []int) []int {

	items_length := len(items)

	if items_length <= 1 {
		return items
	}

	flag := items[rand.Intn(items_length)]

	left := make([]int, 0, items_length)
	middle := make([]int, 0, items_length)
	right := make([]int, 0, items_length)

	for _, item := range items {
		switch {
		case item < flag:
			left = append(left, item)
		case item == flag:
			middle = append(middle, item)
		case item > flag:
			right = append(right, item)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	left = append(left, middle...)
	left = append(left, right...)

	return left
}

func main() {
	fmt.Println(quickSort([]int{1, 6, 2, 9, 10, 8, 0, 5}))
}
