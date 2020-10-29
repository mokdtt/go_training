package main

import (
	"io"
)

//io.Readerを満たした構造体
type LimitedReader struct {
	r     io.Reader
	next  int
	limit int64
}

//Readを持っていなくてはいけない
func (lr *LimitedReader) Read(p []byte) (int, error) {
	n, err := lr.r.Read(p[:lr.limit])
	lr.next += n
	if int64(lr.next) >= lr.limit {
		return n, io.EOF
	}
	return int(n), err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, 0, n}
}
