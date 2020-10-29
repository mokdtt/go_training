package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	c int64
	w io.Writer
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	b.w.Write(p)
	b.c += int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var b ByteCounter
	b.w = w
	return &b, &(b.c)
}

func main() {
	w, c := CountingWriter(os.Stdout)
	w.Write([]byte("hello world\n"))
	fmt.Println(*c)
}
