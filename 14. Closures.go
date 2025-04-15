package main

import "fmt"

// 形式：函数嵌套定义函数，且存在自由变量（介于外部函数和内部函数之间定义的），
// 运行机制：闭包捕获，通过变量逃逸分析捕获到堆上，以获得更长的生命周期
// 意：类似持久化一个函数
// 应用： 状态保持，封装等等（todo）
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
