package subbench

//============================================================
// Table Driven Benchmarks
//============================================================

func Join(parts []string) string {
	out := ""
	for _, p := range parts {
		out += p
	}
	return out
}
