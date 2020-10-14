package main

import (
	"fmt"
)

func main() {
	a := returnOne()
	fmt.Println(a)
}

func returnOne() (i int) {
	defer func() {
		if p := recover(); p != nil {
			i = 1
		}
	}()
	panic("a")
}
