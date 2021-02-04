go build -o bzipper main.go
wc -c < /usr/share/dict/words
./bzipper < /usr/share/dict/words | wc -c
