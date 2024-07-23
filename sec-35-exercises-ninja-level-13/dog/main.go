// Package dog provides a function to convert human years to dog years
package dog

// Years converts human years to dog years
func Years(humanYears int) int {
	return humanYears * 7
}

// YearsTwo converts human years to dog years
func YearsTwo(humanYears int) int {
	count := 0
	for i := 0; i < humanYears; i++ {
		count += 7
	}
	return count
}
