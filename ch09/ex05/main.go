package main

import (
	"fmt"
	"time"
)

func main() {
	ball := make(chan struct{})
	cnt := 0
	go func() {
		ball <- struct{}{}
		for {
			ball <- <-ball
			cnt++
		}
	}()
	go func() {
		for {
			ball <- <-ball
			cnt++
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Printf("%.1f回/sec\n", float64(cnt)/10.)
		return
	}
}
