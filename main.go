package main

import (
	"errors"
	"fmt"
)

func chain(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if n == 0 {
		panic(errors.New("Cannot multiply a number by zero"))
	} else {
		fmt.Println(n)
	}
}

func main() {
	chain(2)
	chain(4)
	chain(0)
	chain(8)
}
