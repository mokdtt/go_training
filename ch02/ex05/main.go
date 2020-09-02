package main

import (
	"fmt"
	"go_training/ch02/ex05/popcountbottom"
	"gopl.io/ch2/popcount"
	"os"
	"strconv"
)

func main() {
	s := os.Args[1]
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("===== Original =====")
	count := popcount.PopCount(uint64(n))
	fmt.Printf("population count = %d\n", count)

	fmt.Println("===== New ======")
	count2 := popcountbottom.PopCount(uint64(n))
	fmt.Printf("population count = %d", count2)
}
