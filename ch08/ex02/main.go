package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	ROOTDIR = "sample_server_root"
)

func handleConn(c net.Conn) {
	defer c.Close()
	//接続が確認できたら
	_, err := io.WriteString(c, "220 sample FTP server ready\n")
	if err != nil {
		return
	}

	var dataConn net.Conn //データコネクション
	currentDir := ROOTDIR //現在のdirectoryを記録

	input := bufio.NewScanner(c)
	for input.Scan() {
		reqStr := input.Text()
		fmt.Println(reqStr)
		reqs := strings.Split(reqStr, " ")
		switch reqs[0] {
		case "USER":
			s := fmt.Sprintf("%sさん,パスワードを入力してください\n", reqs[1])
			io.WriteString(c, "331 "+s)
		case "PASS":
			io.WriteString(c, "230 仮認証OK\n")
		case "SYST":
			io.WriteString(c, "215 UNIX\n")
		case "PORT":
			dataConn, err = setDataconn(reqs[1])
			if err != nil {
				log.Print(err)
				io.WriteString(c, "425 Can't open data connection\n")
				continue
			}
			io.WriteString(c, "200 PORT command successful\n")
		case "LIST":
			io.WriteString(c, "150 File status okay\n")
			err := execLs(reqs, currentDir, dataConn) //とりあえずcurrentDirだけ
			if err != nil {
				log.Print(err)
				io.WriteString(c, "550 Failed\n")
				continue
			}
			io.WriteString(c, "226 Closing data connection\n")
		case "CWD":
			currentDir, err = changeDirectory(reqs[1], currentDir)
			if err != nil {
				io.WriteString(c, "550 Failed\n")
			} else {
				io.WriteString(c, "250 CWD command successful\n")
			}
		case "PWD":
			io.WriteString(c, fmt.Sprintf("257 %s\n", currentDir)) //表示されない
		case "TYPE":
			io.WriteString(c, "200 スルー(本当はすべきでない)\n")
		case "RETR":
			io.WriteString(c, "150 File status okay\n")
			err := getFile(reqs[1], currentDir, dataConn)
			if err != nil {
				log.Print(err)
				io.WriteString(c, "550 Failed\n")
				continue
			}
			io.WriteString(c, "226 Transfer complete: Closing data connection\n")
		case "STOR":
			io.WriteString(c, "150 File status okay\n")
			err := putFile(reqs[1], currentDir, dataConn)
			if err != nil {
				log.Print(err)
				io.WriteString(c, "550 Failed\n")
				continue
			}
			io.WriteString(c, "226 Transfer complete: Closing data connection\n")
		case "QUIT":
			io.WriteString(c, "221 さようなら\n")
			return
		default:
			s := fmt.Sprintf("Not implemented yet: %s\n", reqs[0])
			io.WriteString(c, "504 "+s)
		}
	}
}

func main() {
	fmt.Printf("ROOTDIRを%sとして実装\n", ROOTDIR)
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
