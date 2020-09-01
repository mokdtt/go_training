package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("echoFor")
	s1 := EchoFor(os.Args[1:])
	fmt.Println(s1)
	fmt.Println("echoJoin")
	s2 := EchoJoin(os.Args[1:])
	fmt.Println(s2)
}

func EchoFor(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func EchoJoin(args []string) string {
	s := strings.Join(args, " ")
	return s
}
