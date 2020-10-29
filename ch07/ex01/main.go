package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c += 1
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		*c += 1
	}
	return len(p), nil
}

func main() {
	var c1 WordCounter
	c1.Write([]byte("hello world"))
	fmt.Println(c1)

	var c2 LineCounter
	s := `aaa
bbb
ccc`
	c2.Write([]byte(s))
	fmt.Println(c2)
}
