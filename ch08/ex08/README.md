### 動かし方
- reverb.go動かす
- netcat.go動かす

### 結果
```
❯ go run reverb.go
input:  aaa
input:  bbb
timeout
```

```
❯ go run netcat.go
aaa
         AAA
         aaa
bbb
         BBB
         aaa
         bbb
         bbb
ccc
```
->何も返ってこない(正しい)
