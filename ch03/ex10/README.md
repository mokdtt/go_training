- ベンチマーク関数の実行の結果

- 文字列が長いほど再帰呼び出ししない方が強そう？

```
❯ go test -bench .
goos: darwin
goarch: amd64
pkg: go_training/ch03/ex10
BenchmarkComma9-4                5824801               186 ns/op
BenchmarkCommaOrig9-4            7119708               155 ns/op
BenchmarkComma60-4               1309240               893 ns/op
BenchmarkCommaOrig60-4            609914              1949 ns/op
PASS
ok      go_training/ch03/ex10   6.595s
```
