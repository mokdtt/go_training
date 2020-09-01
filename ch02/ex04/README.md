### 性能比較
- 出力
goos: darwin
goarch: amd64
pkg: go_training/ch02/ex04/popcountshift
BenchmarkPopCount-4        	1000000000	         0.394 ns/op
BenchmarkPopCountShift-4   	21104232	        60.1 ns/op
PASS
ok  	go_training/ch02/ex04/popcountshift	2.261s

- shiftを使うと処理速度も遅くなり，testされる回数も少なくなる
	- loopよりも遅い
