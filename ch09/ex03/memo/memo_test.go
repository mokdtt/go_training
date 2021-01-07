package memo

import (
	"testing"

	"go_training/ch09/ex03/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

//1周目がキャンセルされて，2周目で時間をかけて取得
func TestCancel(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.SequentialCancel(t, m)
}
