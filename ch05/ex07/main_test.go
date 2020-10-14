package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	tests := []string{"http://gopl.io", "https://golang.org"}
	for _, url := range tests {
		descr := fmt.Sprintf("outline(%q)", url)
		out = new(bytes.Buffer)
		if err := outline(url); err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
		got := out.(*bytes.Buffer).String()
		_, err := html.Parse(strings.NewReader(got))
		if err != nil {
			t.Errorf("Parse failed: %v", err)
		}
	}
}
