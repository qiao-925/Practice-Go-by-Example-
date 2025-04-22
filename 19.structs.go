package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name, s.age)

	// 使用同一个堆内存
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	// 深拷贝策略 1. 挨个点名
	sp_deepcopy := person{
		name: s.name,
		age:  s.age,
	}
	fmt.Println("sp_deepcopy", sp_deepcopy)

	sp_deepcopy.name = "peter"
	sp_deepcopy.age = 28

	fmt.Println("sp_deepcopy", sp_deepcopy)
	fmt.Println("s", s)

	// 深拷贝策略 2.json序列化
	sp_deepcopy_byjson := person{}
	byte, error := json.Marshal(s)
	if error != nil {
		log.Fatal(error)
	}
	error = json.Unmarshal(byte, &sp_deepcopy_byjson)
	if error != nil {
		log.Fatal(error)
	}
	sp_deepcopy_byjson.name = "bridge"
	sp_deepcopy_byjson.age = 28
	fmt.Println("sp_deepcopy_byjson", sp_deepcopy_byjson)
	fmt.Println("s", s)

	//------ continue
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
