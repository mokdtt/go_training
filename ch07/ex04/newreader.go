package main

import (
	"io"
)

//io.Readerを満たした構造体
type MyReader struct {
	s string
}

//Readを持っていなくてはいけない
func (m *MyReader) Read(p []byte) (int, error) {
	n := copy(p, m.s) //sをpへコピー
	m.s = m.s[n:]     //読んだ部分を更新
	if len(m.s) == 0 {
		return n, io.EOF
	}
	return n, nil
}

func NewReader(s string) io.Reader {
	return &MyReader{s}
}
