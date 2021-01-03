package intset

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkIntSetAdd(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var s IntSet
		n := rng.Intn(1000000)
		s.Add(n)
	}
}

func BenchmarkMapIntSetAdd(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var s MapIntSet
		n := rng.Intn(1000000)
		s.Add(n)
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	var s1, s2 IntSet
	for i := 0; i < 10000; i++ {
		n := rng.Intn(1000000)
		s1.Add(n)
		s2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.UnionWith(&s2)
	}
}

func BenchmarkMapIntSetUnionWith(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	var s1, s2 MapIntSet
	for i := 0; i < 10000; i++ {
		n := rng.Intn(1000000)
		s1.Add(n)
		s2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.UnionWith(&s2)
	}
}
