// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", Comma(os.Args[i]))
	}
}

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
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
	return buf.String()
}

// comma inserts commas in a non-negative decimal integer string.
func CommaOrig(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return CommaOrig(s[:n-3]) + "," + s[n-3:]
}
