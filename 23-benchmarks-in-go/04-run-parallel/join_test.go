package benchbasics

import "testing"

//=============================================================
// Run Parallel
//=============================================================

func BenchmarkJoin(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		parts := []string{"1", "2", "3"}

		var result string
		for pb.Next() {
			result = Join(parts)
		}
		_ = result
	})
}
