package compileropt

//============================================================
// Compiler Optimizations
//============================================================

func CountDivisible(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			count++
		}
	}
	return count
}
