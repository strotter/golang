package main

import "testing"

func BenchmarkAlloc1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Alloc1()
	}
}

func BenchmarkAlloc2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Alloc2()
	}
}
