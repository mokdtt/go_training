// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package lengthconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// CToF converts a Celsius temperature to Fahrenheit.
func FToM(f Feet) Meter { return Meter(f / 3.2808) }

// FToC converts a Fahrenheit temperature to Celsius.
func MToF(m Meter) Feet { return Feet(m * 3.2808) }

//!-
