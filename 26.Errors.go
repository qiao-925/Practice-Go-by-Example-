package main

import (
	"errors"
	"fmt"
)

// TeaMaker 结构体，封装了茶相关的常量和方法
type TeaMaker struct {
	// 定义与茶相关的常量
	InvalidValue int
	TeaLimit     int
	PowerError   int
}

// NewTeaMaker 创建 TeaMaker 实例
func NewTeaMaker() *TeaMaker {
	return &TeaMaker{
		InvalidValue: 42,
		TeaLimit:     2,
		PowerError:   4,
	}
}

// processValue 函数处理输入值，如果输入值为设定的无效值，则返回错误
func (tm *TeaMaker) processValue(inputValue int) (int, error) {
	if inputValue == tm.InvalidValue {
		return -1, fmt.Errorf("cannot process value %d", tm.InvalidValue)
	}
	return inputValue + 3, nil
}

// ErrNoTea 可用茶叶不足的错误
var ErrNoTea = fmt.Errorf("no more tea available")

// ErrNoPower 无法烧水的错误
var ErrNoPower = fmt.Errorf("cannot boil water")

// prepareTea 函数准备茶，根据输入值返回不同的错误
func (tm *TeaMaker) prepareTea(inputValue int) error {
	if inputValue == tm.TeaLimit {
		return ErrNoTea
	} else if inputValue == tm.PowerError {
		return fmt.Errorf("preparing tea: %w", ErrNoPower)
	}
	return nil
}

func main() {
	// 创建 TeaMaker 实例
	teaMaker := NewTeaMaker()

	// 测试 processValue 函数
	for _, i := range []int{7, teaMaker.InvalidValue} {
		result, err := teaMaker.processValue(i)
		if err != nil {
			fmt.Println("processValue failed:", err)
		} else {
			fmt.Println("processValue worked:", result)
		}
	}

	// 测试 prepareTea 函数
	for i := range 5 {
		err := teaMaker.prepareTea(i)
		if err != nil {
			switch {
			case errors.Is(err, ErrNoTea):
				fmt.Println("We should buy new tea!")
			case errors.Is(err, ErrNoPower):
				fmt.Println("Now it's dark.")
			default:
				fmt.Printf("Unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}
