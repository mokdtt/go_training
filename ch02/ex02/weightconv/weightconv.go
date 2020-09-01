// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package weightconv

import "fmt"

type Pond float64
type Kilogram float64

func (p Pond) String() string     { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func PToK(p Pond) Kilogram { return Kilogram(p / 2.2046) }

// FToC converts a Fahrenheit temperature to Celsius.
func KToP(k Kilogram) Pond { return Pond(k * 2.2046) }

//!-
