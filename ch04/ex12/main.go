package main

import (
	"fmt"
	"os"
)

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string
}

type WordIndex map[string]map[int]bool
type NumIndex map[int]Comic

const usage = `Usage xkcd Command:
	xkcd {QUERY}`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	fmt.Println("Fetch comics ...")
	comicChan, err := getComics()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Done(Fetch)")

	fmt.Println("Indexing ...")
	query, filename := os.Args[1], os.Args[1]
	err = indexing(filename, comicChan)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Done(Indexing)")

	fmt.Printf("search: %s\n", query)
	err = searchWord(query, filename)
	if err != nil {
		os.Exit(1)
	}
}
