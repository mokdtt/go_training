// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name, address, time string
}

func (c *clock) Write(data []byte) (n int, err error) {
	c.time = strings.TrimRight(string(data), "\n")
	return len(data), err
}

func main() {
	clocks, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := range clocks {
		c := &clocks[i]
		conn, err := net.Dial("tcp", c.address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(c, conn)
	}
	for {
		s := ""
		sep := ""
		for _, c := range clocks {
			ss := fmt.Sprintf("%s: %s", c.name, c.time)
			s += sep + ss
			sep = ", "
		}
		fmt.Printf("\r%s", s)
		time.Sleep(1 * time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func parseArgs() ([]clock, error) {
	clocks := []clock{}
	if len(os.Args) == 1 {
		return nil, fmt.Errorf("TZ and address not specified")
	}
	for i := 1; i < len(os.Args); i++ {
		s := strings.Split(os.Args[i], "=")
		if len(s) != 2 {
			return nil, fmt.Errorf("invalid format")
		}
		clocks = append(clocks, clock{s[0], s[1], ""})
	}
	return clocks, nil
}
