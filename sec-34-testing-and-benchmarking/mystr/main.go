package mystr

import "strings"

// Slow version
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// Fast version
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
