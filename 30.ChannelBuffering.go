package main

import "fmt"

func main() {
	msg := make(chan string, 4)

	msg <- "have"
	msg <- "a"
	msg <- "nice"
	msg <- "Tuesday"

	close(msg)

	for s := range msg {
		fmt.Println(s)
	}
}
