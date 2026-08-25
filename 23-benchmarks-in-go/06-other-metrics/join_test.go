package othermetrics

import "testing"

// ============================================================
// Other Metrics
// ============================================================

func BenchmarkJoin(b *testing.B) {
	parts := []string{"1", "2", "3"}

	b.SetBytes(3)

	for b.Loop() {
		Join(parts)
	}

	b.ReportMetric(100, "metric")
}
