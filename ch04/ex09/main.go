// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ファイル名を入力してください")
		os.Exit(1)
	}

	for _, filepath := range os.Args[1:] {
		fmt.Printf("> %s\n", filepath)
		counts := wordfreq(filepath)
		for c, n := range counts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

}

func wordfreq(filepath string) map[string]int {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		return nil
	}
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts
}
