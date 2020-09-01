### 性能比較
- 出力
goos: darwin
goarch: amd64
pkg: go_training/ch02/ex05/popcountbottom
BenchmarkPopCount-4         	1000000000	         0.393 ns/op
BenchmarkPopCountBottom-4   	40290603	        29.2 ns/op
PASS
ok  	go_training/ch02/ex05/popcountbottom	2.083s

- 最下位ビットをクリアしていく方法を使うと処理速度も遅くなり，testされる回数も少なくなる
	- 今回のテストだと2.3のものより遅いが，ビットシフトの方法よりは早い
