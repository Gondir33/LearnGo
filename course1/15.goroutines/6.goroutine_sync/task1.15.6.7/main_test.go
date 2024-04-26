package main

import (
	"sync"
	"testing"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() any { return Person{} },
}

func BenchmarkWithoutPool(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	// benchmark code
}

func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	// benchmark code
}
