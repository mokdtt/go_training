package main

import (
	"fmt"
)

const (
	KB = 1e3
	MB = KB * 1e3
	GB = MB * 1e3
	TB = GB * 1e3
	PB = TB * 1e3
	EB = PB * 1e3
	ZB = EB * 1e3
	YB = ZB * 1e3
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
