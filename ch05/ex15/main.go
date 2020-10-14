package main

import (
	"fmt"
)

func main() {
	vals1 := []int{1, 2, 3, 4, 5}
	vals2 := []int{-5, -4, -3, -2, -1}
	fmt.Println(max(vals1...))
	fmt.Println(max(vals2...))
	fmt.Println(min(vals1...))
	fmt.Println(min(vals2...))
	fmt.Println(max())
	fmt.Println(max2(0, vals1...)) //明示的に必要
	fmt.Println(max2(0, vals2...)) //明示的に必要
	//fmt.Println(max2(vals1...)) //こうは書けない
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	vmax := vals[0]
	for _, val := range vals {
		if val > vmax {
			vmax = val
		}
	}
	return vmax, nil
}

func max2(v0 int, vals ...int) int {
	vmax := v0
	for _, val := range vals {
		if val > vmax {
			vmax = val
		}
	}
	return vmax
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	vmin := vals[0]
	for _, val := range vals {
		if val < vmin {
			vmin = val
		}
	}
	return vmin, nil
}
