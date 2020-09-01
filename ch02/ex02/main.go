// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"go_training/ch02/ex02/lengthconv"
	"go_training/ch02/ex02/weightconv"
	"gopl.io/ch2/tempconv" //versionによっては相対importでは動かない
)

func main() {
	var ts []string
	if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			s := input.Text()
			// 終了条件
			if s == "end" {
				break
			}
			ts = append(ts, s)
		}
	} else {
		ts = os.Args[1:]

	}
	for _, arg := range ts {
		t, err := strconv.ParseFloat(arg, 65)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("===== Temprature Conversion ======")
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Println("===== Length Conversion ======")
		fe := lengthconv.Feet(t)
		m := lengthconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			fe, lengthconv.FToM(fe), m, lengthconv.MToF(m))
		fmt.Println("===== Weight Conversion ======")
		p := weightconv.Pond(t)
		g := weightconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.PToK(p), g, weightconv.KToP(g))

	}
}

//!-
