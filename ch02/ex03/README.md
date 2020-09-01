### 性能比較
- 出力
goos: darwin
goarch: amd64
pkg: go_training/ch02/ex03/popcountloop
BenchmarkPopCount-4       	1000000000	         0.392 ns/op
BenchmarkPopCountLoop-4   	53344231	        22.4 ns/op
PASS
ok  	go_training/ch02/ex03/popcountloop	2.026s

- loopを使うと処理速度も遅くなり，testされる回数も少なくなる
