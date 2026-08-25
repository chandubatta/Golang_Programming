package benchbasics

import "testing"

//============================================================
// Benchmarks Basics
//============================================================
func BenchmarkJoin(b *testing.B) {
	parts := []string{"1", "2", "3"}

	for b.Loop() {
		Join(parts) // --benchtime - 1s
	}
}
