- 実行結果
```
~/go_training/ch09/ex06 master*
❯ ./run.sh
GOMAXPROCS=1
        7.87 real         7.44 user         0.15 sys
GOMAXPROCS=2
        5.18 real         7.94 user         0.14 sys
GOMAXPROCS=3
        4.72 real        10.17 user         0.16 sys
GOMAXPROCS=4
        4.51 real        11.57 user         0.16 sys
GOMAXPROCS=5
        4.38 real        11.51 user         0.16 sys
GOMAXPROCS=6
        4.30 real        11.48 user         0.15 sys
GOMAXPROCS=7
        4.28 real        11.56 user         0.15 sys
```
- 自分のPC
	- 2コア4スレッド
