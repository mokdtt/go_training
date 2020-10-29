package main

import (
	"fmt"

	"go_training/ch07/ex03/treesort"
)

func main() {
	fmt.Println("===== example 1 =====")
	input := []int{4, 3, 2, 1, 5}
	var root *treesort.Tree
	for _, v := range input {
		root = treesort.Add(root, v)
	}
	fmt.Println(root)
	fmt.Println("===== example 2 =====")
	input2 := []int{5, 6, 3, 4, 1, 8, 9, 2, 7}
	var root2 *treesort.Tree
	for _, v := range input2 {
		root2 = treesort.Add(root2, v)
	}
	fmt.Println(root2)
}
