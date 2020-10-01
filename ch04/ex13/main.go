package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `Usage poster Command:
	poster {MOVIE TITLE}`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	title := strings.Join(os.Args[1:], "+")
	movie, err := getMovie(title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(movie.Title, movie.Year)
	fmt.Println(movie.Poster)
	err = getPoster(movie.Poster, strings.Join(os.Args[1:], "-"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
	}
}
