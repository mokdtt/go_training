package main

import (
	"fmt"
)

func pipeline(n int) (chan struct{}, chan struct{}) {
	first := make(chan struct{})
	last := make(chan struct{})
	prev := first
	for i := 0; i < n; i++ {
		out := make(chan struct{})
		go func(in chan struct{}, out chan struct{}) {
			for {
				<-in
				out <- struct{}{}
			}
		}(prev, out)
		prev = out
		last = out
	}
	return first, last
}

func main() {
	for i := 1; ; i++ {
		fmt.Println(i)
		first, last := pipeline(i)
		first <- struct{}{}
		<-last
	}
}
