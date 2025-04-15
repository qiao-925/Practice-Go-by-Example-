package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, "\n")

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// 函数参数为slice时，需要加...
	sum(nums...)
}
