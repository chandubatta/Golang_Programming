package timerorder

import (
	"testing"
	"testing/synctest"
	"time"
)

// ============================================================
// Timer Ordering
// ============================================================

func TestOrdering(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var order []string

		for _, name := range []string{"a", "b", "c"} {
			time.AfterFunc(time.Second, func() {
				order = append(order, name)
			})
		}

		synctest.Sleep(time.Second)
		t.Log("order:", order)
	})
}
