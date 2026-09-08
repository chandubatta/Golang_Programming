package theproblem

import (
	"context"
	"time"
)

// ============================================================
// The Problem
// ============================================================

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
