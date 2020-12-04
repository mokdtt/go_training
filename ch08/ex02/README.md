# FTP serverの実装
- serverを動かす
```
$ run_server.sh
```
- clientを動かす
```
$ cd workspace
$ ftp 127.0.0.1 8000
```

### 参考
- [FTP（File Transfer Protocol）～前編](https://www.atmarkit.co.jp/ait/articles/0107/17/news002.html)
	- FTPの基本概念がわかる
- [別表1　FTPのコマンド一覧](https://www.atmarkit.co.jp/fnetwork/rensai/netpro11/ftp-command.html)
	- 内部的に何を実装すれば良いかがわかる
- [別表1　FTPのレスポンスコード一覧](https://www.atmarkit.co.jp/fnetwork/rensai/netpro10/ftp-responsecode.html)
	- どういうコードを返すべきかがわかる
