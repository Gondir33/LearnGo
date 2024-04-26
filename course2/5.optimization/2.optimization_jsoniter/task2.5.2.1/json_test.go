package main

import (
	"os"
	"testing"
)

func BenchmarkMarshalJson(b *testing.B) {
	bytes, _ := os.ReadFile("test.josn")
	test := &CheckJson{}
	var r Welcome
	test.Unmarshal(bytes, &r)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		test.Marshal(r)
	}
}

func BenchmarkUnmarshalJson(b *testing.B) {
	bytes, _ := os.ReadFile("test.josn")
	test := &CheckJson{}
	b.ResetTimer()
	b.ReportAllocs()
	var r Welcome

	for i := 0; i < b.N; i++ {
		test.Unmarshal(bytes, &r)
	}
}

func BenchmarkMarshalJsoniter(b *testing.B) {
	bytes, _ := os.ReadFile("test.josn")
	test := &CheckJsoniter{}
	var r Welcome
	test.Unmarshal(bytes, &r)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		test.Marshal(r)
	}
}

func BenchmarkUnmarshalJsoniter(b *testing.B) {
	bytes, _ := os.ReadFile("test.josn")
	test := &CheckJsoniter{}
	b.ResetTimer()
	b.ReportAllocs()
	var r Welcome

	for i := 0; i < b.N; i++ {
		test.Unmarshal(bytes, &r)
	}
}

func BenchmarkMarshalEasyjson(b *testing.B) {
	bytes, _ := os.ReadFile("test.josn")
	test := &CheckEasyjson{}
	var r Welcome
	test.Unmarshal(bytes, &r)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		test.Marshal(r)
	}
}

func BenchmarkUnmarshalEasyjson(b *testing.B) {
	bytes, _ := os.ReadFile("test.josn")
	test := &CheckEasyjson{}
	b.ResetTimer()
	b.ReportAllocs()
	var r Welcome

	for i := 0; i < b.N; i++ {
		test.Unmarshal(bytes, &r)
	}
}
