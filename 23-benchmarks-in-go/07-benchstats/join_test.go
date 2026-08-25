package benchbasics

import "testing"

// ============================================================
// Benchstats
// ============================================================
func BenchmarkJoin(b *testing.B) {
	parts := make([]string, 10000)
	for i := range 10000 {
		parts[i] = "chunk"
	}

	for b.Loop() {
		Join(parts)
	}
}
