package compileropt

import "testing"

//============================================================
// Compiler Optimizations
//============================================================

func BenchmarkCountDivisible(b *testing.B) {
	for b.Loop() {
		CountDivisible(1000)
	}
}

func Wrap() {
	CountDivisible(1000)
}

func BenchmarkCountDivisibleInline(b *testing.B) {
	for b.Loop() {
		Wrap()
	}
}
