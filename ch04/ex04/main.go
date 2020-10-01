// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 1)
	fmt.Println(s)
	rotate(s, 2)
	fmt.Println(s)
	rotate(s, 5)
	fmt.Println(s)
	rotate(s, 6)
	fmt.Println(s)
}

// Rotates in place the slice s to the left by n steps
func rotate(s []int, n int) {
	if len(s) == 0 || n >= len(s) {
		fmt.Println("Warning: no rotate (n >= len(s))")
		return
	}

	n = n % len(s)

	ss := make([]int, len(s))

	copy(ss, s)
	copy(s, ss[n:])
	copy(s[len(s)-n:], ss[0:n])
}
