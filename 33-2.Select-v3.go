package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

/*
*
使用超时的context,
适用场景：
- 外部API可能存在的超时情况
- 数据库查询超时
*/
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		time.Sleep(time.Second)

		c1 <- "one"
	}()

	go func() {
		// 超时和任务完成切换
		time.Sleep(2 * time.Second)
		// time.Sleep(5 * time.Second)
		c2 <- "two"
	}()

	var receivedCount int32 = 0
	totalMsg := 2

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done:", ctx.Err())
			return
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			atomic.AddInt32(&receivedCount, 1)

		case msg2 := <-c2:
			fmt.Println("received", msg2)
			atomic.AddInt32(&receivedCount, 1)
		}

		if atomic.LoadInt32(&receivedCount) == int32(totalMsg) {
			fmt.Println("All messages received, exiting...")
			return
		}
	}

}
