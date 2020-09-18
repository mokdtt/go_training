package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", CommaFloat(os.Args[i]))
	}
}

func CommaFloat(s string) string {
	stmp := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s, stmp = s[:i], s[i+1:]
			break
		}
	}
	n := len(s)
	if n <= 3 {
		if stmp == "" {
			return s
		}
		return s + "." + stmp
	}

	rem := n % 3
	n = n + (n-1)/3

	var buf bytes.Buffer
	buf.Grow(n)
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if (i+1)%3 == rem && (i+1) != len(s) {
			buf.WriteString(",")
		}
	}
	if stmp == "" {
		return buf.String()
	}
	return buf.String() + "." + stmp
}
