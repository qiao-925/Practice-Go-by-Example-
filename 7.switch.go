package main

// import 在编写时/后统一处理
import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")
	// 省略break
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It`s a weekend")
	default:
		fmt.Println("It`s a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")

	default:
		fmt.Println("It`s after noon")
	}

	whatAmI := func(i interface{}) {
		switch tp := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("Match failed.type:%T", tp)
		}
	}

	whatAmI(true)
	whatAmI(32)
	whatAmI("String")

}
