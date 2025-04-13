package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 13

	fmt.Println("map", m)

	v1 := m["key1"]
	fmt.Println("v1", v1)

	// 越界默认返回该类型零值
	v3 := m["key3"]
	fmt.Println("v3", v3)

	fmt.Println("len", len(m))

	delete(m, "key2")
	fmt.Println("map", m)

	clear(m)
	fmt.Println("map", m)

	value, exist := m["key2"]
	fmt.Println("value", value)
	fmt.Println("exist", exist)

	newMap := map[string]int{"foo": 1, "bar": 2}

	newMap2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(newMap, newMap2) {
		fmt.Println("newMap == newMap2")
	}

}
