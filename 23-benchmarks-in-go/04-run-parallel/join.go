package benchbasics

//============================================================
// Run Parallel
//============================================================

func Join(parts []string) string {
	out := ""
	for _, p := range parts {
		out += p
	}
	return out
}
