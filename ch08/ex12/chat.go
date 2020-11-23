// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type cliData struct {
	client   chan<- string
	username string
}

//type client chan<- string // an outgoing message channel

var (
	entering = make(chan cliData)
	leaving  = make(chan cliData)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[cliData]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.client <- msg
			}

		case cli := <-entering:
			users := []string{}
			for c := range clients {
				users = append(users, c.username)
			}
			msg := "online user: " + strings.Join(users, ",")
			cli.client <- msg
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.client)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cliData{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cliData{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
