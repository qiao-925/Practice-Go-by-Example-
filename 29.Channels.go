package main

import "fmt"

func main() {
	msgChannel := make(chan string)

	go func() {
		msgChannel <- "ping"

	}()

	fmt.Println(<-msgChannel)

}
