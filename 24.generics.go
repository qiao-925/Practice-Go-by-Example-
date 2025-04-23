package main

import "fmt"

// ~[]E 是类型约束，表示 S 的底层类型 (underlying type) 必须是 []E。
// ~ 符号表示 "近似" (approximately)，意味着 S 可以是 []E 本身，
// 也可以是基于 []E 的自定义类型。
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) AllElements() []T {
	var elems []T
	for e := list.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")

	list := List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)
	fmt.Println("list:", list.AllElements())
}
