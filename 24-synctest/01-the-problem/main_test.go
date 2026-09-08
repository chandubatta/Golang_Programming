package theproblem

import (
	"context"
	"errors"
	"testing"
	"time"
)

// ============================================================
// The Problem
// ============================================================

func TestRetry(t *testing.T) {
	start := time.Now()
	var offsets []time.Duration
	calls := 0

	Retry(context.Background(), 3, time.Second, func() error {
		offsets = append(offsets, time.Since(start))
		calls++

		if calls < 3 {
			return errors.New("Boom")
		}

		return nil
	})

	want := []time.Duration{0, time.Second, 3 * time.Second}
	const slop = 50 * time.Millisecond

	for i, got := range offsets {
		if got < want[i] || got > want[i]+slop {
			t.Errorf("attempt %d at %v, want %v (+/- %v)", i, got, want[i], slop)
		}
	}
}

func TestRetryFast(t *testing.T) {
	start := time.Now()
	var offsets []time.Duration
	calls := 0

	Retry(context.Background(), 3, time.Millisecond, func() error {
		offsets = append(offsets, time.Since(start))
		calls++

		if calls < 3 {
			return errors.New("Boom")
		}

		return nil
	})

	want := []time.Duration{0, time.Millisecond, 3 * time.Millisecond}
	const slop = 1 * time.Millisecond

	for i, got := range offsets {
		if got < want[i] || got > want[i]+slop {
			t.Errorf("attempt %d at %v, want %v (+/- %v)", i, got, want[i], slop)
		}
	}
}
