https://moz.com/top500よりcsvをダウンロード．
list化する

## 実行
```
$ cat urls.txt | xargs go run main.go
```

## あるウェブサイトが応答しない時
- localhost:3000で何も実行せずに試してみる
```
Get http://localhost:3000: dial tcp [::1]:3000: connect: connection refused
```
となる
