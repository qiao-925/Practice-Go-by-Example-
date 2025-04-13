package main

import "fmt"

func main() {
	// 下标为容量
	var a [5]int
	fmt.Println("init array a ", a)

	// 下标为索引
	a[4] = 100
	fmt.Println("set a[4] ", a)
	fmt.Println("get a[4] ", a[4])
	fmt.Println("len :", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init and set", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("init and set", b)

	// 选择性赋值
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d:", twoD)

}
