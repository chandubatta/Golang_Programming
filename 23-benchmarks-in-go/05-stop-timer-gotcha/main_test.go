package stoptimer

import (
	"math/rand"
	"slices"
	"testing"
)

//============================================================
// StopTimer
//============================================================

func BenchmarkSort(b *testing.B) {
	r := rand.New(rand.NewSource(1))

	for b.Loop() {

		b.StopTimer()
		s := make([]int, 1000)
		for i := range s {
			s[i] = r.Int()
		}
		b.StartTimer()

		slices.Sort(s)
	}
}

func BenchmarkSortCopy(b *testing.B) {
	r := rand.New(rand.NewSource(1))

	s := make([]int, 1000)
	for i := range s {
		s[i] = r.Int()
	}

	sCopy := make([]int, 1000)

	for b.Loop() {
		copy(sCopy, s)
		slices.Sort(sCopy)
	}
}
