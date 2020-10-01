package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var f1 = flag.Bool("sha384", false, "Use SHA384")
var f2 = flag.Bool("sha512", false, "Use SHA512")

func main() {
	flag.Parse()
	err := *f1 && *f2
	if err {
		fmt.Fprintf(os.Stderr, "flagは一つしか指定できません")
		os.Exit(1)
	}
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	c := ""
	if *f1 {
		c = fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
	} else if *f2 {
		c = fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
	} else {
		c = fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
	}
	fmt.Println(c)
}
