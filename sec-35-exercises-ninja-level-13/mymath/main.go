// Package mymath provides mathematical functions
package mymath

import "sort"

// CenteredAvg calculates the average of a slice of integers, removing the highest and lowest values
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1 : len(xi)-1] // remove first and last elements

	sum := 0
	for _, v := range xi {
		sum += v
	}

	return float64(sum) / float64(len(xi))
}
