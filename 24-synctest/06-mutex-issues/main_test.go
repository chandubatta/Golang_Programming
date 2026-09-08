package mutexhang

import (
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

// ============================================================
// Mutex Issues
// ============================================================

func TestMutexHangs(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var mu sync.Mutex
		mu.Lock()

		go func() {
			synctest.Sleep(time.Second)
			mu.Unlock()
		}()

		mu.Lock()
	})
}
