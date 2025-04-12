package main

import "fmt"

/*
*
  - 多变量同时初始化赋值
  - 省略var写法“：”
    手上功夫
  - 细节
  - 熟悉度
  - 直觉，反应（知道和做到的过程）
  - 细节感知
*/
func main() {
	var a = "initinal"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

}
