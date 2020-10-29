package main

import (
	"fmt"

	"go_training/ch07/ex13/eval"
)

func main() {
	expr, _ := eval.Parse("sqrt(100)-2*4")
	fmt.Println(expr)
}
