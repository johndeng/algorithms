// Merge sort is an O(n log n) comparison-based sorting algorithm.

package main

import "fmt"

// Merge slice
func merge(left, right []int) []int {

	left_length := len(left)
	right_length := len(right)
	total_length := left_length + right_length

	result := make([]int, total_length)
	n, m := 0, 0

	for n < len(left) && m < len(right) {

		if left[n] <= right[m] {
			result[n+m] = left[n]
			n++
		} else {
			result[n+m] = right[m]
			m++
		}
	}

	for n < len(left) {
		result[n+m] = left[n]
		n++
	}
	for m < len(right) {
		result[n+m] = right[m]
		m++
	}

	return result

}

// Split slice and merge it
func MergeSort(items []int) []int {

	if len(items) <= 1 {
		return items
	}

	middle := int(len(items) / 2)

	left := MergeSort(items[:middle])
	right := MergeSort(items[middle:])
	return merge(left, right)

}

func main() {
	fmt.Println(MergeSort([]int{1, 3, 5, 11, 9, 4, 6, 2, 3}))
}
