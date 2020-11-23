package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

var maxDepth = flag.Int("depth", 3, "max depth")
var tokens = make(chan struct{}, 20)

type depthURL struct {
	links []string
	depth int
}

func crawl(url string, depth int) *depthURL {
	if depth >= *maxDepth {
		return &depthURL{nil, depth + 1}
	}
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return &depthURL{list, depth + 1}
}

func main() {
	flag.Parse()
	worklist := make(chan *depthURL)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- &depthURL{flag.Args(), 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, depth int) {
					worklist <- crawl(link, depth)
				}(link, list.depth)
			}
		}
	}
}
