package main

import (
	"fmt"
	"my_sort" // 导入自定义包
)

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := my_sort.IntArray(data)
	my_sort.Sort(a)
	// if !my_sort.IsSorted(a) {
	// 	panic("fails")
	// }
	fmt.Printf("The sorted array is: %v\n", data)
}
