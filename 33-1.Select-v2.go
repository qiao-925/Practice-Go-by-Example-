package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// UploadFile 模拟上传文件的函数
func UploadFile(ctx context.Context, filename string) error {
	fmt.Printf("开始上传文件: %s\n", filename)
	defer fmt.Printf("文件上传完成: %s\n", filename)

	// 模拟上传过程
	for i := 0; i <= 100; i += 10 {
		select {
		case <-ctx.Done():
			fmt.Printf("上传取消: %s, 进度: %d%%\n", filename, i)
			return ctx.Err() // 返回 Context 的错误
		default:
			// 模拟上传耗时
			sleepTime := time.Duration(rand.Intn(500)) * time.Millisecond
			time.Sleep(sleepTime)
			fmt.Printf("上传进度: %s, %d%%\n", filename, i)
		}
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保退出时取消 Context

	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}

	// 启动多个 Goroutine 上传文件
	for _, filename := range filenames {
		go func(filename string) {
			err := UploadFile(ctx, filename)
			if err != nil {
				fmt.Printf("文件上传失败: %s, 错误: %v\n", filename, err)
			}
		}(filename)
	}

	// 生产环境中，通常通过api来中止上传，这意味着需要存储上传时的状态以及context,以便通过context取消任务

	// 模拟用户在 2 秒后取消上传
	time.Sleep(2 * time.Second)
	fmt.Println("取消上传...")
	cancel() // 取消 Context，通知所有 Goroutine 停止上传

	// 等待一段时间，让 Goroutine 完成清理操作
	time.Sleep(3 * time.Second)
	fmt.Println("上传程序退出")
}
