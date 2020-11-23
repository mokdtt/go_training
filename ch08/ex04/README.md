### 動かし方
- reverb.go動かす
- netcat.go動かす
- Ctrl-Dを押して，受信を続けることを確認する

### 結果
- rever_orig.goを裏で動かしていた時にCtrl-Dを挟む
```
❯ go run ./netcat.go
aa
         AA
         aa
         aa
bb
         BB
2020/11/21 14:01:02 done
```

- rever_orig.goを裏で動かしていた時にCtrl-Dを挟む
```
❯ go run ./netcat.go
aa
         AA
         aa
         aa
bb
         BB
^D       bb
         bb
2020/11/21 14:02:23 done
```
