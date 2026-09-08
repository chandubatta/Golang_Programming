package bubbledthings

import (
	"testing"
	"testing/synctest"
)

// ============================================================
// Bubbled Things Can't Leave
// ============================================================

func TestChannelEscapes(t *testing.T) {
	var ch chan int

	synctest.Test(t, func(t *testing.T) {
		ch = make(chan int)
	})

	ch <- 1
}
