package main

import (
	"bufio"
	"fmt"
	"os"

	"go_training/ch07/ex15/eval"
)

func main() {
	fmt.Printf("式を入力してください: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	inputStr := input.Text()
	if inputStr == "" {
		os.Exit(1)
	}
	expr, err := eval.Parse(inputStr)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println(expr.Eval(eval.Env{}))
}
