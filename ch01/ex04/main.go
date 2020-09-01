package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)          //文字の出現回数
	flgs := make(map[string]map[string]int) //["a"]["in1.txt"]のような二重構造
	files := os.Args[1:]
	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, flgs)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for key, _ := range flgs[line] {
				fmt.Printf(key + "\t")
			}
			fmt.Println("")
		}
	}
}

func countLines(f *os.File, filename string,
	counts map[string]int, flgs map[string]map[string]int) error {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if _, ok := flgs[input.Text()]; !ok {
			flgs[input.Text()] = make(map[string]int)
		}
		flgs[input.Text()][filename] = 1
	}
	// NOTE: ignoring potential errors from input.Err()
	return nil
}
