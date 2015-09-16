// Quick sort algorithms is invented by Hoare
package main

import (
	"fmt"
)

func quickSort(items []int, begin, end int) {

	flag := begin
	left := begin
	right := end

	if begin >= end {
		return
	}

	for begin < end {

		// 从右到左, 找到比标记值大的数
		for begin < end {

			if items[end] < items[flag] {
				items[begin], items[end] = items[end], items[begin]
				// 重置起点
				flag = end
				break
			} else {
				end -= 1
				continue
			}
		}

		// 从左到右, 找到比标记值小的数
		for begin < end {

			if items[begin] > items[flag] {
				items[begin], items[end] = items[end], items[begin]
				// 重置起点
				flag = begin
				break
			} else {
				begin += 1
				continue
			}
		}
	}

	// 划分比 flag 所在值小的列表
	quickSort(items, left, flag-1)

	// 划分比 flag 所在值大的列表
	quickSort(items, flag+1, right)

}

func main() {
	items := []int{4, 5, 7, 2, 3, 1}
	quickSort(items, 0, len(items)-1)
	fmt.Println(items)
}
