package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	r := []byte(s)
	fmt.Println(string(r))
	reverse(r)
	fmt.Println(string(r))
}

func reverse(bs []byte) {
	sizeList := []int{}
	for i := 0; i < len(bs); {
		_, size := utf8.DecodeRune(bs[i:])
		sizeList = append(sizeList, size)
		i += size
	}
	//全体を逆順に
	reverseBytes(bs)

	//個別に逆順
	loc := 0
	for i := 0; i < len(sizeList); i++ {
		size := sizeList[len(sizeList)-1-i]
		reverseBytes(bs[loc : loc+size])
		loc += size
	}
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

}
