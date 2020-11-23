// Fetch saves the contents of a URL into a local file.
package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var done = make(chan struct{})
var wg sync.WaitGroup

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func fetch(url string, ch chan<- *http.Response) {
	defer wg.Done()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	ch <- resp
}

func main() {
	ch := make(chan *http.Response)
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, ch)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	resp := <-ch
	defer resp.Body.Close()
	fmt.Println(resp.Request.URL)
	fmt.Println(resp.Header)
}
