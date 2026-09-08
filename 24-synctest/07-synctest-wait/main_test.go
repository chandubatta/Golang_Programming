package wait

import (
	"testing"
	"testing/synctest"
)

// ============================================================
// synctest.Wait
// ============================================================

func TestUploadNoWait(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		u := NewUploader()
		defer u.Close()

		u.Upload("a")
		u.Upload("b")

		synctest.Wait()

		if got := len(u.Uploaded()); got != 2 {
			t.Errorf("uploaded %d, want 2", got)
		}
	})
}
