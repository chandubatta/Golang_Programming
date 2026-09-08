package poisonedselect

import (
	"testing"
	"testing/synctest"
	"time"
)

// ============================================================
// Poisoned Select
// ============================================================

var shutdown = make(chan struct{})

func TestPoisonedSelect(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		work := make(chan int)

		go func() {
			synctest.Sleep(time.Second)
			work <- 1
		}()

		select {
		case <-work:
		case <-shutdown:
		}
	})
}
