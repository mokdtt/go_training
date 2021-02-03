// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Sample struct {
		B    bool
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
	}
	input := Sample{true, 1.0, 1.0, 1.1 + 2.2i, 3.3 + 4.4i}
	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var sample Sample
	if err := Unmarshal(data, &sample); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", sample)

	// Check equality.
	if !reflect.DeepEqual(sample, input) {
		t.Fatal("not equal")
	}
}

func TestInterface(t *testing.T) {
	type Sample struct {
		I interface{}
	}
	input := Sample{[]int{1, 2, 3}}
	data, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var sample Sample
	if err := Unmarshal(data, &sample); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", sample)

	// Check equality.
	if !reflect.DeepEqual(sample, input) {
		t.Fatal("not equal")
	}
}
