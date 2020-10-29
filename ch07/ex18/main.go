// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (n *Element) String() string {
	b := &bytes.Buffer{}
	forEachNode(n, b, 0)
	return b.String()
}

func forEachNode(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case CharData:
		s := strings.TrimSpace(string(n))
		fmt.Fprintf(w, "%*s%s\n", depth*2, "", s)
	case *Element:
		for _, c := range n.Children {
			fmt.Fprintf(w, "%*s<%s>\n", depth*2, "", n.Type.Local)
			forEachNode(c, w, depth+1)
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Type.Local)
		}
	}
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	stack := []*Element{&Element{xml.Name{"root", ""}, nil, []Node{}}}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			next := &Element{tok.Name, tok.Attr, []Node{}}
			tmp := stack[len(stack)-1]
			tmp.Children = append(tmp.Children, next)
			stack = append(stack, next)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			tmp := stack[len(stack)-1]
			tmp.Children = append(tmp.Children, CharData(tok))
		}
	}
	fmt.Println(stack[0])
}
