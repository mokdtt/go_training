package main

import (
	"testing"
)

func BenchmarkEchoFor(b *testing.B) {
	letters := []string{"a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		EchoFor(letters)
	}
}

func BenchmarkEchoJoin(b *testing.B) {
	letters := []string{"a", "b", "c", "d"}
	for i := 0; i < b.N; i++ {
		EchoJoin(letters)
	}
}
