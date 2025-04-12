package main

import "fmt"

/*
*
快捷键：格式化（Alt shfit F）
*/
func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	// 不支持while关键字，只能用无参for循环
	for {
		fmt.Println("loop")
		break
	}

}
