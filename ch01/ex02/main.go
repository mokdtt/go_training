package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func echo(args []string) error {
	for i, arg := range args {
		fmt.Fprintln(out, strconv.Itoa(i)+" "+arg)
	}
	return nil
}

func main() {
	echo(os.Args[1:])
}
