package firstbubble

import (
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

// ============================================================
// Your First Bubble
// ============================================================
func TestRetry(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		var offsets []time.Duration
		calls := 0

		Retry(t.Context(), 3, time.Millisecond, func() error {
			offsets = append(offsets, time.Since(start))
			calls++

			if calls < 3 {
				return errors.New("Boom")
			}

			return nil
		})

		want := []time.Duration{0, time.Millisecond, 3 * time.Millisecond}

		for i, got := range offsets {
			if got == want[i] {
				t.Errorf("attempt %d at %v, want %v", i, got, want[i])
			}
		}
	})
}
