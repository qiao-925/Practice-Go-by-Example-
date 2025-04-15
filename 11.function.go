package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// fmt.Printf可用于多变量的格式化输出
func main() {
	a, b := vals()
	fmt.Printf("a: %v,b:%v\n", a, b)

	_, c := vals()
	fmt.Println(c)
}
