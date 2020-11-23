package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	sCh := make(chan string)

	go func() { // 文字列をバッファなしへ送る
		for input.Scan() {
			sCh <- input.Text()
		}
	}()
L:
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("timeout")
			break L
		case s := <-sCh:
			fmt.Println("input: ", s)
			go echo(c, s, 1*time.Second)
		}
	}
	close(sCh)
	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
