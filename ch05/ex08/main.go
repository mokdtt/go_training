package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

const usage = `sample usage
	$ go run main.go url id`

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]
	err := outline(url, id)
	if err != nil {
		fmt.Printf("outline(%s) failed: %v\n", url, err)
		os.Exit(1)
	}
}

func outline(url string, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	n := ElementByID(doc, id)
	if n == nil {
		return fmt.Errorf("can't find ID")
	}
	fmt.Printf("<%s", n.Data)
	for _, attr := range n.Attr {
		fmt.Printf(" %s=%q", attr.Key, attr.Val)
	}
	fmt.Printf(">\n")
	return nil
}

func forEachNode(n *html.Node, pre func(n *html.Node) bool) {
	if pre != nil {
		if pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
	return
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var nodeID *html.Node
	findID := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					nodeID = n
					return true
				}
			}
		}
		return false
	}
	forEachNode(doc, findID)
	return nodeID
}
