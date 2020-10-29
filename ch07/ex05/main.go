package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	example := "abcdefghijk"
	reader := LimitReader(strings.NewReader(example), 3)
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	fmt.Println(buf.String())
}
