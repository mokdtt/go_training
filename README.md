# go_training

## 実行
```
$ go run main.go
もしくは
$ ./run.sh
```

## テスト
```
$ go test
$ go test -v
```

## 自作パッケージ管理
- go.modを作成
```
$ go mod init go_training
```

- importするとき
```
import (
	"go_training/ex**/ch**/**"
)
```

- 自作packageをinstall
```
$ go install go_training/ch**/ex**/**
```
