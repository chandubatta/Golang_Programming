package subbench

import "testing"

//=============================================================
// Table Driven Benchmarks
//=============================================================

func getStringSlice(size int) []string {
	s := make([]string, size)
	for i := range size {
		s[i] = "chunk"
	}

	return s
}

func BenchmarkJoin(b *testing.B) {
	benches := []struct {
		name string
		slice []string
	}{
		{
			name: "10",
			slice: getStringSlice(10),
		},
		{
			name: "100",
			slice: getStringSlice(100),
		},
		{
			name: "1000",
			slice: getStringSlice(1000),
		},
	}

	for _, bench := range benches {
		b.Run(bench.name, func(b *testing.B) {
			for b.Loop() {
				Join(bench.slice)
			}
		})
	}
}
