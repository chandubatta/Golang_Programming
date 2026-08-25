package othermetrics

//============================================================
// Other Metrics
//============================================================

func Join(parts []string) string {
	out := ""
	for _, p := range parts {
		out += p
	}
	return out
}
