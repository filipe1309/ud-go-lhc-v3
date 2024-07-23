// Package acdc provides some functions to sum numbers
package acdc

// Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
