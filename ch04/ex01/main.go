package main

import (
	"crypto/sha256"
	"fmt"
	"gopl.io/ch2/popcount"
)

func DiffCount(b1 byte, b2 byte) (cnt int) {
	i := uint64(^(b1 ^ b2))
	//fmt.Printf("%b, %b, %b\n", b1, b2, i)
	return popcount.PopCount(i)
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	cnt := 0
	for i := 0; i < len(c1); i++ {
		cnt += DiffCount(c1[i], c2[i])
	}
	fmt.Println(cnt)
}
