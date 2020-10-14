package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ElementsByTagName: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("===== img =====")
	image := ElementsByTagName(doc, "img")
	for _, n := range image {
		fmt.Println(n.Data)
	}
	fmt.Printf("\n===== h1 h2 h3 h4 =====\n")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, n := range headings {
		fmt.Println(n.Data)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var listNode []*html.Node
	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, tag := range name {
				if tag == n.Data {
					listNode = append(listNode, n)
				}
			}
		}
	}

	forEachNode(doc, pre)
	return listNode
}

func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}
