package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func savePage(u string) error {
	parsedURL, err := url.Parse(u)
	urlHost := parsedURL.Host
	// 一旦ホストが同じかチェック
	if okHost != urlHost {
		return fmt.Errorf("allowed host %s, but got %s", okHost, urlHost)
	}
	// pathとファイル名を決定する
	filedir := filepath.Join("tmp", urlHost, parsedURL.Path)
	filename := filepath.Join(filedir, "index.html")
	// Get
	resp, err := http.Get(u)
	if err != nil {
		return fmt.Errorf("getting %s: %s", u, resp.Status)
	}
	defer resp.Body.Close()
	// ディレクトリ作成
	err = os.MkdirAll(filedir, 0777)
	if err != nil {
		return err
	}
	// ファイル書き込み
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, body, 0664); err != nil {
		return err
	}
	fmt.Println("saved")
	return nil
}

func crawl(url string) []string {
	fmt.Println(url)
	err := savePage(url)
	if err != nil {
		fmt.Printf("can't save page - %s: %v\n", url, err)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var okHost string

const usage = `go run main.go URL 
`

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	parsed, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	okHost = parsed.Host
	breadthFirst(crawl, os.Args[1:])
}
