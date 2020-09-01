// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// FToC converts a Celsius.to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC converts a Kelvin to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

//!-
