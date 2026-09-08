package firstbubble

import (
	"context"
	"time"
)

// ============================================================
// Your First Bubble
// ============================================================

// Retry calls do until it succeeds, backing off exponentially between
// attempts. It gives up early if ctx is cancelled.
func Retry(ctx context.Context, attempts int, base time.Duration, do func() error) error {
	var err error
	for i := range attempts {
		if err = do(); err == nil {
			return nil
		}

		if i == attempts-1 {
			break
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(base << i):
		}
	}
	return err
}
