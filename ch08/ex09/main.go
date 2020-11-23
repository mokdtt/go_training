package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type fileSizes struct {
	id   int
	size int64
}

func main() {
	// ...determine roots...
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizesByRoot := make(chan fileSizes)
	var n sync.WaitGroup
	for id, root := range roots {
		n.Add(1)
		go walkDir(id, root, &n, fileSizesByRoot)
	}
	go func() {
		n.Wait()
		close(fileSizesByRoot)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles := make(map[int]int64)
	nbytes := make(map[int]int64)
loop:
	for {
		select {
		case fs, ok := <-fileSizesByRoot:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[fs.id]++
			nbytes[fs.id] += fs.size
		case <-tick:
			printDiskUsage(nfiles, nbytes, roots)
		}
	}

	printDiskUsage(nfiles, nbytes, roots) // final totals
}

func printDiskUsage(nfiles, nbytes map[int]int64, roots []string) {
	for i, name := range roots {
		fmt.Printf("%s: %d files  %.1f GB\n", name, nfiles[i], float64(nbytes[i])/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(id int, dir string, n *sync.WaitGroup, fileSizesByRoot chan<- fileSizes) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(id, subdir, n, fileSizesByRoot)
		} else {
			fileSizesByRoot <- fileSizes{id, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
