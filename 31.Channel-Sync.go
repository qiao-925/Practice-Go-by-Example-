package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")

	time.Sleep(time.Second)

	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 2)
	go worker(done)

	/**

	按照这段代码的逻辑，main 函数确实有一个任务在等待 worker
	函数执行完毕后才能继续执行。具体来说，main 函数的任务就是等待 worker
	函数完成，然后程序才能正常退出。
	*/
	fmt.Println(<-done)
}
