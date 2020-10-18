package main

import (
	"fmt"

	"go_training/ch06/ex01/intset"
)

func main() {
	var x intset.IntSet
	fmt.Println("x:", x.Len(), x.String())
	x.Add(1)
	fmt.Println("x:", x.Len(), x.String())
	x.Add(144)
	fmt.Println("x:", x.Len(), x.String())
	x.Add(9)
	fmt.Println("x:", x.Len(), x.String())
	x.Remove(1)
	fmt.Println("x:", x.Len(), x.String())
	y := x.Copy()
	fmt.Println("y:", y.Len(), y.String())
	x.Clear()
	fmt.Println("x:", x.Len(), x.String())
}
