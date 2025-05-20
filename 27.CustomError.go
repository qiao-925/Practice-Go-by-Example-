package main

import (
	"errors"
	"fmt"
)

type ArgError struct {
	Arg int
	Msg string
}

func (e *ArgError) Error() string {

	return fmt.Sprintf("Arg: %d, %s ", e.Arg, e.Msg)

}
func NewArgError(arg int, msg string) error {
	return &ArgError{Arg: arg, Msg: msg}
}

func processArg(arg int) (int, error) {
	if arg == 42 {
		return -1, NewArgError(arg, "Don`t use 42")
	}

	return arg + 3, nil
}

func main() {
	_, err := processArg(42)

	var argErr *ArgError
	if errors.As(err, &argErr) {
		fmt.Println("Arg:", argErr.Arg)
		fmt.Println("Msg:", argErr.Msg)
	} else {
		fmt.Println("Error is not an ArgError")
	}
}
