package main

import (
	"fmt"
	"slices"
)

func main() {
	var slice []string
	fmt.Println("uninit", slice, slice == nil, len(slice) == 0)

	slice = make([]string, 3)
	fmt.Println("emp", slice, "len", len(slice), "cap", cap(slice))

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"

	fmt.Println("print slice:", slice)
	fmt.Println("slice [2]", slice[2])

	fmt.Println("len", len(slice))

	// 返回值再赋值
	slice = append(slice, "d", "e")

	fmt.Println("len", len(slice))

	// ==================
	copyTest := make([]string, len(slice))

	copy(copyTest, slice)

	fmt.Println("copyTest", copyTest)

	cutTestslice := slice[2:4]
	fmt.Println("cutTestslice", cutTestslice)

	cutTestslice = slice[:3]
	fmt.Println("cutTestslice", cutTestslice)

	cutTestslice = slice[3:]
	fmt.Println("cutTestslice", cutTestslice)

	compareTest1 := []string{"a", "b"}
	compareTest2 := []string{"a", "b"}

	if slices.Equal(compareTest1, compareTest2) {
		fmt.Println("compareTest1 == compareTest2")
	}

	// 嵌套
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
