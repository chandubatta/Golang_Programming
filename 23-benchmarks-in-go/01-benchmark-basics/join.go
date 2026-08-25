package benchbasics

//============================================================
// Benchmarks Basics
//============================================================

func Join(parts []string) string {
	out := ""
	for _, p := range parts {
		out += p
	}
	return out
}
